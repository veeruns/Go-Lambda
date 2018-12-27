package main

import "fmt"

type code int
const (
  CRITICAL code = iota
  WARNING
  OK
)
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

func (re result) Validate() (bool,string){
  if(len(re.outputDesc) == 0 || len(re.outputString) == 0) || re.outputCode > 3 || re.outputCode < 0 ){
    return false,"Something Wrong with result set"
  } else {
    return true,"All OK"
  }
}

var Result result
