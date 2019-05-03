package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func basicAuth() string {
	var username string = "snowensemble"
	var passwd string = "KbxYDhpQJNdDS1IuQnt!$2uy!QVp1Px6Q0mRtHpU*"

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://vzbuilders.service-now.com/api/now/table/u_ops_request_management", nil)

	req.SetBasicAuth(username, passwd)
	query := req.URL.Query()
	query.Add("sysparm_query", "u_business_service%3D8434f15b0fc0cfc0abe3590be1050e86^state=1")
	query.Add("sysparam_limit", "10")
	req.URL.RawQuery = query.Encode()

	//fmt.Println(string(requestDump))
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

func main() {
	//fmt.Println("requesting...")
	S := basicAuth()
	fmt.Println(S)
}
