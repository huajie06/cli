package main

import (
	"coreutils/cmd"
	"flag"
	"fmt"
	"os"
)

func main() {
	// var cmdString []string
	// cmdString = []string{"-n"}
	// cmd.Uname(cmdString)
	// cmd.Arch()

	// m()

	// cmd.Base64("helloworld today is a good day")
	cmd.Base64("YQ==")
}

func m() {
	var CommandLine = flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	if err := CommandLine.Parse(os.Args[1:]); err != nil {
		fmt.Println("error is ", err)
	}

	args := CommandLine.Args()
	// pgrm := strings.ToLower(args[0])
	pgrm := args[0]
	params := args[1:]

	switch pgrm {
	case "arch":
		cmd.Arch()
	case "uname":
		cmd.Uname(params)
	default:
		_, err := fmt.Fprintf(os.Stderr, "no such program: %s\n", pgrm)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
