package rokulib

import (
	"testing"
)

func TestReadconfig(t *testing.T) {
	var c Config
	op := readconfig(&c, "conf", "Server")
	if c.Certdir != "/opt/httpsServer/conf/certs" {
		t.Error("Expected /opt/httpsServer/conf/certs")
	}
	if op == false {
		t.Fatal("It should always be true")
	}
}

func TestWrongdir(t *testing.T) {
	var c Config
	op := readconfig(&c, "/opt/httpsServer/conf", "Server")
	if op == true {
		t.Fatal("There should not be anything there")
	}
}
