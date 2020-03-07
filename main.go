package main

import (
	"coreutils/cmd"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	var CommandLine = flag.NewFlagSet(os.Args[1], flag.ExitOnError)

	if err := CommandLine.Parse(os.Args[1:]); err != nil {
		fmt.Println("error is ", err)
	}

	_ = CommandLine.Bool("l", true, "some usage")
	if err := CommandLine.Parse(os.Args[1:]); err != nil {
		fmt.Println("error is ", err)
	}

	args := CommandLine.Args()
	pgrm := strings.ToLower(args[0])
	// params := args[1:]
	// fmt.Println(args)

	// fmt.Println("prgam is: ", pgrm)
	// fmt.Println("options are: ", params)

	switch pgrm {
	case "arch":
		cmd.Arch()
	case "ls":
		fmt.Println("ls")
	default:
		fmt.Println("no such program")
	}
}
