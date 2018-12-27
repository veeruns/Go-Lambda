package main

import "fmt"

type result struct {
	outputDesc   string
	outputCode   int
	outputString string
}

func (re result) PrintResult() {
	fmt.Printf("OutputCode is %d\n", re.outputCode)
	fmt.Printf("OutputString is %s\n", re.outputString)
	fmt.Printf("OutputDesc is %s\n", re.outputDesc)

}

var Result result
