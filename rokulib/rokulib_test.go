package rokulib

import (
	"testing"
)

func Testreadconfig(t *testing.T) {
	var c Config
	readconfig(&c)
	if c.Certdir != "/opt/httpsServer/conf/certs" {
		t.Error("Exepcted /opt/httpsServer/conf/certs")
	}
}
