package main

import "fmt"

var disk Result

func main() {
	disk.outputString = "Test"
	fmt.Printf("%s\n", disk.outputString)
}
