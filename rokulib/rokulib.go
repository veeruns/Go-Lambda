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

var datachan chan *HttpResponse
var signal chan string

//ch := make(chan *HttpResponse, 1)

//PowerOff Roku box
//InitLib initializes channel
func InitLib() {
	datachan := make(chan *HttpResponse, 1)
	signal := make(chan string)
	//start workerpool
	go workerpool()
}

//PowerOff function sends poweroff to roku
func PowerOff(hostname string) bool {
	//	var buff bytes.Buffer

	var somedata HttpResponse
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/keypress/poweroff")
	somedata.url = url.String()
	datachan <- &somedata
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

func asynchttp() string {

	//var resps *HttpResponse
	var buff bytes.Buffer
	var url string
	var recievedata *HttpResponse
	recievedata = <-datachan
	url = recievedata.url

	fmt.Printf("Fetching Roku URL %s\n", url)
	resp, err := http.Post(url, "", &buff)
	defer resp.Body.Close()

	return "stuffed"
}

func workerpool() {
	for j := range datachan {
		fmt.Println("Started reading from data channel")
	}
	var resps *HttpResponse
	for {
		select {
		case r := <-datachan:
			fmt.Printf("%s was fetched\n", r.url)
			resps = r
			return resps
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}
