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

// bitwise operation demo
// https://play.golang.org/p/VeLCx_4orSW

const mapEncoder = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// const NoPadding rune = -1

func newEnc(encoder string) *encoding {
	e := new(encoding)
	e.padChar = '='
	copy(e.encode[:], encoder) // a good way to assign value
	// when you have a slice make([]int{}, 10), you can dynamically pop the value
	// by using copy(dest, src []Type), slice[:] <=> string
	// dest=[0 0 0 0 0 0 0] src=[1 2 3] => [1 2 3 0 0 0 0]

	for i := 0; i < len(e.decodeMap); i++ {
		e.decodeMap[i] = 0xFF // 0xff = 255, fill with 255
	}

	// string[index] => byte => uinit8
	// "abc"[0] => 97
	for i := 0; i < len(encoder); i++ {
		e.decodeMap[encoder[i]] = byte(i)
	}
	return e
}

func (enc *encoding) Encode(src []byte) []byte {
	if len(src) == 0 {
		return []byte{}
	}

	dst := make([]byte, (len(src)+2)/3*4)

	// 1 byte = 8 bits, so everything 3`letters`/bytes = 24 bits => 4 new `blocks`
	// everything 3 letters => 4 new letters, remaining will fill accordingly
	// so the length will have to be `(n + 2) / 3 * 4`

	di, si := 0, 0
	n := (len(src) / 3) * 3

	for si < n {
		// Convert 3x 8bit source bytes into 4 bytes
		val := uint(src[si+0])<<16 | uint(src[si+1])<<8 | uint(src[si+2])
		// if src[si+0] = "a", "a"
		// 01100001 => 011000010000000000000000, add sixteen zeros
		// 01100001 =>         0110000100000000, add 8 zeros
		// 01100001 =>                 01100001, stays the same
		//          => 011000010110000101100001 => 68371921 in decimal
		// use |(or), to get an overall number

		dst[di+0] = enc.encode[val>>18&0x3F]
		dst[di+1] = enc.encode[val>>12&0x3F]
		dst[di+2] = enc.encode[val>>6&0x3F]
		dst[di+3] = enc.encode[val&0x3F]
		// 0x3f = 63 = 111111
		// 011000                   & 111111 => 011000 => 24 => enc.encode[24]
		// 011000010110             & 111111 => 010110 => 22 =>
		// 011000010110000101       & 111111 => 000101 => 5  =>
		// 011000010110000101100001 & 111111 => 100001 => 33 =>
		si += 3
		di += 4
	}

	remain := len(src) - si
	if remain == 0 {
		return dst
	}

	switch remain {
	case 2:

		val := uint(src[si+0])<<10 | uint(src[si+1])<<2
		dst[di+0] = enc.encode[val>>12&0x3f]
		dst[di+1] = enc.encode[val>>6&0x3F]
		dst[di+2] = enc.encode[val&0x3F]
		dst[di+3] = byte(enc.padChar)
	case 1:
		val := uint(src[si+0]) << 4
		dst[di+0] = enc.encode[val>>6&0x3f]
		dst[di+1] = enc.encode[val&0x3F]
		dst[di+2] = byte(enc.padChar)
		dst[di+3] = byte(enc.padChar)
	}
	return dst
}

func Base64(s string) {
	e := newEnc(mapEncoder)
	r := e.Encode([]byte(s))
	fmt.Println(string(r))
}
