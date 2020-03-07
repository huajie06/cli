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

	a := flag.String("name", "default value", "description")
	flag.Parse()

	// ProgName := os.Args[0]
	// Usage := os.Args[1]
	// Args := os.Args[2:]
	//fmt.Println(ProgName)
	//fmt.Println(Usage, Args)

	//flags := make([]string)

	// for _, v := range Args {
	// 	fmt.Println(v)
	// }

	fmt.Println(*a)
}
