package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	f := fsm.NewFSM(
		"locked",
		fsm.Events{
			{Name: "unlock", Src: []string{"locked"}, Dst: "unlocked"},
			{Name: "open", Src: []string{"unlocked"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "unlocked"},
			{Name: "lock", Src: []string{"unlocked"}, Dst: "locked"},
		},
		fsm.Callbacks{},
	)

	f.Event("unlock")
	fmt.Println(f.Current())

	f.Event("open")
	fmt.Println(f.Current())
	f.Event("close")
	f.Event("lock")
	err := f.Event("Magic")
	if err != nil {
		fmt.Printf("There is an error %s\n", err.Error())
	}

}
