package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

type trafficlight struct {
	light string
	FSM   *fsm.FSM
}

func Newtrafficlight(light string) *trafficlight {
	l := &trafficlight{
		light: light,
	}

}
func main() {

}
