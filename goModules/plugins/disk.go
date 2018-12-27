package main

import (
	"github.com/veeruns/Go-Lambda/goModules/groucho"
)

func main() {
	var disk groucho.Result
	disk.OutputString = "Test"

	disk.OutputCode = groucho.OK
	disk.OutputDesc = "Test result"
	op, ops := disk.Validate()
	if op == true {
		disk.PrintResult()
	}
}
