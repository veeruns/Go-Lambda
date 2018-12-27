package groucho

import "fmt"

type code int

const (
	CRITICAL code = iota
	WARNING
	OK
)

type Result struct {
	OutputDesc   string
	OutputCode   int
	OutputString string
}

func (re Result) PrintResult() {
	fmt.Printf("OutputCode is %d\n", re.outputCode)
	fmt.Printf("OutputString is %s\n", re.outputString)
	fmt.Printf("OutputDesc is %s\n", re.outputDesc)

}

func (re Result) Validate() (bool, string) {
	if len(re.outputDesc) == 0 || len(re.outputString) == 0 || re.outputCode > 3 || re.outputCode < 0 {
		return false, "Something Wrong with result set"
	}
	return true, "All OK"

}
