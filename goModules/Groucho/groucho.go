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
	OutputCode   code
	OutputString string
}

var AllResults []Result

func (re Result) PrintResult() {
	fmt.Printf("OutputCode is %d\n", re.OutputCode)
	fmt.Printf("OutputString is %s\n", re.OutputString)
	fmt.Printf("OutputDesc is %s\n", re.OutputDesc)

}

func (re Result) Validate() (bool, string) {
	if len(re.OutputDesc) == 0 || len(re.OutputString) == 0 || re.OutputCode > 3 || re.OutputCode < 0 {
		return false, "Something Wrong with result set"
	}
	return true, "All OK"

}

func (re *Result) AppendResults() (bool, string) {
	vop, _ := re.Validate()
	if vop != true {
		return false, "Validation Failure"
	}
	AllResults = append(AllResults, re)
}
