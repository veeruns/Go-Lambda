package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	u := "veeru"
	p := "$R33r@m$$January2019"

	//bouncerRawURL := "https://by.bouncer.login.yahoo.com/login/"
	bouncerRawURL := "http://localhost/login"

	values := make(url.Values)
	values.Add("id", u)
	values.Add("pass_word", p)
	values.Add("action", "login")

	v := values.Encode()
	req, err := http.NewRequest("POST", bouncerRawURL, strings.NewReader(v))
	if err != nil {
		panic(err)
	}

	// Add headers
	req.Header.Add("Accept-Encoding", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", strconv.Itoa(len(v)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Host", "by.bouncer.login.yahoo.com")
	req.Header.Set("User-Agent", "Go HTTP")

	data, err := httputil.DumpRequest(req, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("****REQ****\n%s\n", data)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	data, err = httputil.DumpResponse(resp, false)
	if err != nil {
		panic(err)
	}

	fmt.Printf("****RESP****\n%s\n", data)
}
