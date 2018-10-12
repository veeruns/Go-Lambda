package main

import (
	"fmt"
	"io/ioutil"
	"syscall"
)

func main() {
	if syscall.Getuid() == 0 {
		fmt.Printf("You are running as root, Dropping your privileges")
		err := syscall.Setuid(65534)
		if err != nil {
			fmt.Printf(err.Error())
		} else {

		}

	}
}

func filetostring(filename string) (string, error) {
	var output []byte
	var err error
	output, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	} else {
		return string(output), nil
	}
}
