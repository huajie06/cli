package cmd

import (
	"fmt"
	"os"
	"strings"
)

func Basename(s []string) {
	if len(s) > 1 {
		fmt.Fprintf(os.Stderr, "too many inputs, please only use 1 input")
		return
	}

	if len(s) == 0 {
		fmt.Fprintf(os.Stderr, "Please enter a basename")
		return
	}

	fp := s[0]

	parts := strings.Split(fp, "/")
	fmt.Fprintf(os.Stdout, parts[len(parts)-1])
}
