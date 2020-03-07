package main

import (
	"coreutils/cmd"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var cmdString []string
	cmdString = []string{"abc", "bcd"}
	// cmdString = []string{"-n"}
	//cmd.Uname(cmdString)
	cmd.Arch()

	fmt.Println(cmdString)
	//m()
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
	case "ls":
		fmt.Println("ls")
	case "uname":
		cmd.Uname(params)
	default:
		_, err := fmt.Fprintf(os.Stderr, "no such program: %s\n", pgrm)
		if err != nil {
			log.Fatal(err)
		}
	}
}
