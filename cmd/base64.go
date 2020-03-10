package cmd

import "fmt"

// in shell if you do
// echo "a" | base64
// it by default include a newline character \n

type encoding struct {
	encode    [64]byte
	decodeMap [256]byte
	padChar   rune
}

const mapEncoder = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const NoPadding rune = -1

func newEnc(encoder string) *encoding {
	e := new(encoding)
	e.padChar = '='
	copy(e.encode[:], encoder) // a good way to assign value
	// when you have a slice make([]int{}, 10), you can dynamically pop the value
	// by using copy(dest, src []Type), slice[:] <=> string
	// dest=[0 0 0 0 0 0 0] src=[1 2 3] => [1 2 3 0 0 0 0]

	for i := 0; i < len(e.decodeMap); i++ {
		e.decodeMap[i] = 0xFF // 0xff = 255
	}

	// string[index] => byte => uinit8
	// "abc"[0] => 97
	for i := 0; i < len(encoder); i++ {
		e.decodeMap[encoder[i]] = byte(i)
	}
	return e
}

func (enc *encoding) Encode(dst, src []byte) {
	if len(src) == 0 {
		return
	}

	//_ = enc.encode

	di, si := 0, 0
	n := (len(src) / 3) * 3

	for si < n {
		// Convert 3x 8bit source bytes into 4 bytes
		val := uint(src[si+0])<<16 | uint(src[si+1])<<8 | uint(src[si+2])
		// if src[si+0] = "a", "a"
		// 1100001 => 11000010000000000000000, add sixteen zeros
		// 1100001 => 110000100000000, add 8 zeros
		// 1100001 => 1100001, stays the same
		// use |(or), to get an overall number

		dst[di+0] = enc.encode[val>>18&0x3F]
		dst[di+1] = enc.encode[val>>12&0x3F]
		dst[di+2] = enc.encode[val>>6&0x3F]
		dst[di+3] = enc.encode[val&0x3F]

		si += 3
		di += 4
	}

	remain := len(src) - si
	if remain == 0 {
		return
	}
	// Add the remaining small block
	val := uint(src[si+0]) << 16
	if remain == 2 {
		val |= uint(src[si+1]) << 8
	}

	dst[di+0] = enc.encode[val>>18&0x3F]
	dst[di+1] = enc.encode[val>>12&0x3F]

	switch remain {
	case 2:
		dst[di+2] = enc.encode[val>>6&0x3F]
		if enc.padChar != NoPadding {
			dst[di+3] = byte(enc.padChar)
		}
	case 1:
		if enc.padChar != NoPadding {
			dst[di+2] = byte(enc.padChar)
			dst[di+3] = byte(enc.padChar)
		}
	}
}

func (enc *encoding) EncodedLen(n int) int {
	if enc.padChar == NoPadding {
		return (n*8 + 5) / 6 // minimum # chars at 6 bits per char
	}
	return (n + 2) / 3 * 4 // minimum # 4-char quanta, 3 bytes each
}

func Base64(s string) {
	e := newEnc(mapEncoder)
	buf := make([]byte, e.EncodedLen(len([]byte(s))))
	e.Encode(buf, []byte(s))

	fmt.Println(string(buf))
}
