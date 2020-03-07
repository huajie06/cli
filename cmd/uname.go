package cmd

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/sys/unix"
)

type empty struct{}

// uname only accept 1 flag

func Uname(s []string) {

	if len(s) > 1 {
		fmt.Println("Too many flags")
		return
	}
	unixName := unix.Utsname{}
	if err := unix.Uname(&unixName); err != nil {
		log.Println(err)
	}

	sysName := bytes.Trim(unixName.Sysname[:], "\x00")
	nodeName := bytes.Trim(unixName.Nodename[:], "\x00")
	release := bytes.Trim(unixName.Release[:], "\x00")
	version := bytes.Trim(unixName.Version[:], "\x00")
	machineName := bytes.Trim(unixName.Machine[:], "\x00")

	flags := map[string][]byte{
		"-a": bytes.Join([][]byte{sysName, nodeName, version, release, machineName}, []byte("\n")),
		"-s": sysName,
		"-n": nodeName,
		"-v": version,
		"-r": release,
		"-m": machineName,
	}

	for _, v := range s {
		if _, ok := flags[v]; !(ok) {
			fmt.Println("cmd not existed", v)
			return
		}

	}

	fmt.Printf("%s%s", flags[s[0]], "\n")
}
