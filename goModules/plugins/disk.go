package main

import (
	"github.com/veeruns/Go-Lambda/goModules/groucho"
)

func main() {
	var disk groucho.Result
	disk.OutputString = "Test"
	disk.RegisterID.Name = "DiskChecker"
	disk.RegisterID.Version = "1.0"
	disk.RegisterID.CodePath = "github.com/veeruns/Go-Lambda/goModules/plugins"
	disk.OutputCode = groucho.OK
	disk.OutputDesc = "Test result"
	op, _ := disk.Validate()
	if op == true {
		disk.PrintResult()
	}
}
