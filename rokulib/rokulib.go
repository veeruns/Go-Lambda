package rokulib

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type HttpResponse struct {
	url  string
	resp *http.Response
	err  error
}

//var datachan chan *HttpResponse
var signal chan string

//ch := make(chan *HttpResponse, 1)

//PowerOff Roku box
//InitLib initializes channel
func InitLib() {
	//	datachan := make(chan *HttpResponse, 1)
	signal = make(chan string)
	//start workerpool
	go workerpool()
}

//PowerOff function sends poweroff to roku
func PowerOff(hostname string) bool {
	//	var buff bytes.Buffer
	//	var somedata HttpResponse
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/keypress/poweroff")
	signal <- url.String()
	return true
}

func PowerOn(hostname string) bool {
	var buff bytes.Buffer
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/keypress/poweron")

	resp, err := http.Post(url.String(), "", &buff)
	if err != nil {
		fmt.Printf("That did not work as intended %s, code %d\n", err.Error(), resp.StatusCode)
		return false
	} else {
		//	fmt.Printf("The return string is %s\n", resp.Body)
		return true
	}
}

func workerpool() {
	for {
		fmt.Println("[Rokulib] Started reading from data channel")
		select {
		case msg := <-signal:
			//	var resp *http.Response
			fmt.Printf("[Rokulib]  Recieved to post %s\n ", msg)
			var buff bytes.Buffer
			//resp, err := http.Post(msg, "", &buff)
			time.Sleep(time.Millisecond * 2500)

			/*	if err != nil {
				fmt.Printf("[Rokulib] Error from Roku %s\n", err.Error())
			} else {*/
			fmt.Printf("[Rokulib] Response code from Roku %d\n", resp.StatusCode)
		//	}
		default:
			fmt.Printf("[Rokulib] Nothing recieved yet")

		}
	}
	//	var resps *HttpResponse
}
