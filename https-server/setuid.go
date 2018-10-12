package main

import (
	"fmt"
	"io/ioutil"
	"syscall"
)

func main() {
	if syscall.Getuid() == 0 {
		fmt.Printf("You are running as root, Dropping your privileges")
    toprint,ferr := filetostring("/etc/letsencrypt/live/veeruns.raghavanonline.com/README")
    if(ferr == nil){
      fmt.Printf("README Doc is %s\n",toprint)
    }else{}
      fmt.Printf(ferr.Error())
    }
		err := syscall.Setuid(65534)
		if err != nil {
			fmt.Printf(err.Error())
      _,nerr := filetostring("/etc/letsencrypt/live/veeruns.raghavanonline.com/README")
      if(nerr == nil){
        fmt.Printf("Something wrong it should not happen")
      } else{
        fmt.Printf("This is correct %s\n",nerr.Error())
      }
		} else {


		}

	}
}

func filetostring(filename string) (string, error) {
	var output []byte
	var err error
	output, err = ioutil.ReadFile(filename)
	return string(output), err

}
