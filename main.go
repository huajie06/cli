package main

import (
	"flag"
	"fmt"
	"os"
)

// go run main.go
// ls => a b c
// ls -l=true but vertical
// ls -a=true => a b c .d
// ls -al=true
// ls -a -l

// dat := []string{"a", "b", "c", ".d"}

// cmd := "ls"

// options = []string{"l", "a"}
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

	fmt.Println(args)

}
