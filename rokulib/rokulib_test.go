package rokulib

import (
	"testing"
)

func Testreadconfig(t *testing.T) {
	var c Config
	op := readconfig(&c)
	if c.Certdir != "/opt/httpsServer/conf/cert" {
		t.Error("Exepcted /opt/httpsServer/conf/certs")
	}
	if op == false {
		t.Fatal("It should always be true")
	}
}
