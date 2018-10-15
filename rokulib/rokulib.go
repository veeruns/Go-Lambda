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
func PowerOff(hostname string) bool {
	//	var buff bytes.Buffer
	datachan := make(chan *HttpResponse, 1)
	var somedata HttpResponse
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/keypress/poweroff")
	somedata.url = url.String()
	datachan <- &somedata
	results := asynchttp()
	if results == "stuffed" {
		fmt.Printf("stuffed into channel\n")

	}
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
	url = datachan.url
	go func(url string) {
		fmt.Printf("Fetching Roku URL %s\n", url)
		resp, err := http.Post(url, "", &buff)
		resp.Body.Close()
		datachan <- &HttpResponse{url, resp, err}
	}(url)

	return "stuffed"
}

func getresponses() *HttpResponse {
	var resps *HttpResponse
	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			resps = r
			return resps
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}
