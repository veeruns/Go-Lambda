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
	op := readconfig(&c, "/opt/httpsServer/confs", "Servers")
	fmt.Printf("Output is %v\n", op)
	if c.Certdir == "/opt/httpsServer/conf/certs" {
		t.Error("Expected /opt/httpsServer/conf/certs, there should not be anything here")
	}
	if op == true {
		t.Fatal("There should not be anything there")
	}
}

func TestPowerOff(t *testing.T) {
	InitLib()
	op := PowerOff("10.0.0.3:85")
	if op != true {
		t.Error("Has to return true")
	}
}

func TestPowerOn(t *testing.T) {
	InitLib()
	op := PowerOn("10.0.0.3:85")
	if op != true {
		t.Error("Has to return true")
	}
}
