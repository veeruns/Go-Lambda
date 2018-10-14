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

//PowerOff Roku box
func PowerOff(hostname string) bool {
	var buff bytes.Buffer
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/keypress/poweroff")
	results := asynchttp(url.String())

	if results.err != nil {
		fmt.Printf("That did not work as intended %s\n", results.err.Error())
		return false
	} else {
		fmt.Printf("The return string is %s\n", results.resp.Status)
		return true
	}
}

func PowerOn(hostname string) bool {
	var buff bytes.Buffer
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/keypress/poweron")

	resp, err := http.Post(url.String(), "", &buff)
	if err != nil {
		fmt.Printf("That did not work as intended %s\n", err.Error())
		return false
	} else {
		fmt.Printf("The return string is %s\n", resp.Body)
		return true
	}
}

func asynchttp(url string) *HttpResponse {
	ch := make(chan *HttpResponse, 1)
	resps := *HttpResponse
	var buff bytes.Buffer
	go func(url string) {
		fmt.Printf("Fetching Roku URL %s\n", url)
		resp, err := http.Post(url, "", &buff)
		resp.Body.Close()
		ch <- &HttpResponse{url, resp, err}
	}(url)
	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			resps = append(resps, r)
			if len(responses) == len(urls) {
				return resps
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}

	return resps
}
