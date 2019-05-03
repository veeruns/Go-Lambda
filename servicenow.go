package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func basicAuth() string {
	var username string = "snowensemble"
	var passwd string = "KbxYDhpQJNdDS1IuQnt!$2uy!QVp1Px6Q0mRtHpU"
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://vzbuilders.service-now.com/api/now/table/u_ops_request_management", nil)
	req.SetBasicAuth(username, passwd)
	query := req.URL.Query()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

func main() {
	fmt.Println("requesting...")
	S := basicAuth()
	fmt.Println(S)
}
