package groucho

import (
	"fmt"
)

type code int

type pluginid struct {
	Name     string
	Version  string
	CodePath string
}

const (
	CRITICAL code = iota
	WARNING
	OK
)

type Result struct {
	RegisterID   pluginid
	OutputDesc   string
	OutputCode   code
	OutputString string
}

var inside Result

type allresults []Result

var AllResults allresults

func (re Result) PrintResult() {
	fmt.Printf("OutputCode is %d\n", re.OutputCode)
	fmt.Printf("OutputString is %s\n", re.OutputString)
	fmt.Printf("OutputDesc is %s\n", re.OutputDesc)

}

func (re Result) Validate() (bool, string) {
	if len(re.OutputDesc) == 0 || len(re.OutputString) == 0 || re.OutputCode > 3 || re.OutputCode < 0 || len(re.RegisterID.Name) == 0 || len(re.RegisterID.Version) == 0 {
		return false, "Something Wrong with result set"
	}
	return true, "All OK"

}

func (re allresults) AppendResults(input Result) (bool, string) {
	AllResults = append(AllResults, re)
	return true, "Successfully Appended"
}

//PrintAllResults prints all the data in memory
func (re allresults) PrintAllResults() {
	/*	vop, _ := re.Validate()
		if vop != true {
			function, file, line, _ := runtime.Caller(1)
			op := fmt.Sprintf("Validation failure at %s %s %d", file, runtime.FuncForPC(function).Name(), line)

			return errors.New(op)
		}*/
	for _, v := range re {
		fmt.Printf("%s %s %d\n", v.OutputString, v.OutputDesc, v.OutputCode)
	}
	return nil
}
