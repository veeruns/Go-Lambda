package rokulib

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	var c Config
	op := readconfig(&c, "conf", "Server")
	if c.Certdir != "/opt/httpsServer/conf/certs" {
		t.Error("Expected /opt/httpsServer/conf/certs")
	}
	if op == false {
		t.Fatal("It should always be true")
	}
}

func TestWrongDir(t *testing.T) {
	var c Config
	op := readconfig(&c, "/opt/httpsServer/conf", "Server")
	fmt.Printf("Output is %v\n", op)
	if op == true {
		t.Fatal("There should not be anything there")
	}
}
