package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	simplepb "github.com/veeruns/Go-Lambda/Protocol-Buf/simple"
)

func main() {
	sm := doSimpleStuff()

	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("READ From file", sm2)

}

func writeToFile(fname string, pb proto.Message) error {
	op, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Some error happened ", err)
	}
	if err := ioutil.WriteFile(fname, op, 0644); err != nil {
		log.Fatalln("Cannot write to file ", err)
		return err
	}
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	read, _ := ioutil.ReadFile(fname)
	proto.Unmarshal(read, pb)
	return nil
}
func doSimpleStuff() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Magic Magic Magic",
		SampleList: []int32{1, 2, 3, 4, 5},
	}
	//fmt.Println(sm.GetName())
	return &sm
}
