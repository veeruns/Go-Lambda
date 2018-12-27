package main

import (
	"fmt"

	"github.com/veeruns/Go-Lambda/goModules/groucho"
)

func main() {
	type disk groucho.Result
	disk.outputString = "Test"

	fmt.Printf("%s\n", disk.outputString)
}
