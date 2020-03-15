package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type empty struct{}

func Cat(s []string) {

	e := empty{}
	if len(s) == 0 {
		fmt.Fprintln(os.Stderr, "Please enter a file")
		return
	}

	fmt.Fprintln(os.Stdout, s)

	// -e = display $ at the end of each line
	// -n = display line number
	// -t = display tabs
	flags := map[string]empty{"-e": e, "-n": e, "-t": e}

	// if cat -d, meaning there's a `-`, it will try to parse the flag
	flag, fname := s[0], s[0]
	if strings.Contains(flag, "-") {
		if _, ok := flags[fname]; !ok {
			fmt.Fprintln(os.Stderr, "Flag is not supported.")
			return
		}
		fname = s[1]
	}

	f, err := os.Open(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	switch flag {
	case "-e":
		continue
	case "-n":
		continue
	case "-t":
		continue
	}

	fmt.Fprintf(os.Stdout, buf.String())
	fmt.Println(flag)
}

func displayDollar(b btyes.Buffer) {
	//TODO: add $ at the end of each line
	return
}

func displayLineNum(b bytes.Buffer) {
	return
}

func displayTab(b bytes.Buffer) {
	return
}
