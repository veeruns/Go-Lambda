package main

import (
	"fmt"

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

	var d2 groucho.Result
	d2.OutputString = "test2"
	d2.OutputDesc = "Test2 op"
	d2.OutputCode = groucho.WARNING
	d2.RegisterID.Name = "DiskChecker"
	d2.RegisterID.Version = "2.0"

	var allr groucho.AllResults
	allr.AppendResults(disk)
	fmt.Printf("Done Appeneding")
	allr.PrintAllResults()
	allr.AppendResults(d2)

	op, _ := disk.Validate()
	if op == true {
		disk.PrintResult()
	}
}
