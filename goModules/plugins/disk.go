package main

import (
	"fmt"

	"github.com/veeruns/Go-Lambda/goModules/groucho"
)

var disk groucho.Result

func main() {
	disk.outputString = "Test"

	fmt.Printf("%s\n", disk.outputString)
}
