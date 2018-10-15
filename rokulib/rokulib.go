package rokulib

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

//HTTPResponse a more complex type
type HTTPResponse struct {
	url  string
	resp *http.Response
	err  error
}

type app struct {
	ID          int    `xml:"id,attr"`
	Version     string `xml:"version,attr"`
	Channeltype string `xml:"type,attr"`
	App         string `xml:"app"`
}

type apps struct {
	XMLName xml.Name `xml:"apps"`
	App     []app    `xml:"app"`
}

//var datachan chan *HttpResponse
var signal chan string

//ch := make(chan *HttpResponse, 1)

//InitLib initializes channel
func InitLib() {
	//	datachan := make(chan *HttpResponse, 1)
	signal = make(chan string)
	readchannels()
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

//PowerOn function powers on roku
func PowerOn(hostname string) bool {
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/keypress/poweron")
	signal <- url.String()
	return true
}

func LaunchChannel(hostname string, channelid int) bool {
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/launch/")
	url.WriteString(strconv.Itoa(channelid))
	signal <- url.String()
	return true

}

func workerpool() {
	for {
		fmt.Println("[Rokulib] Started reading from data channel")
		select {
		case msg := <-signal:
			var resp *http.Response
			fmt.Printf("[Rokulib]  Recieved to post %s\n ", msg)
			var buff bytes.Buffer
			resp, err := http.Post(msg, "", &buff)
			//time.Sleep(time.Millisecond * 2500)
			if err != nil {
				fmt.Printf("[Rokulib] Error from Roku %s\n", err.Error())
			} else {
				fmt.Printf("[Rokulib] Response code from Roku %d\n", resp.StatusCode)
			}
		case <-time.After(30 * time.Second):
			fmt.Printf("[Rokulib] Nothing recieved yet")

		}
	}
	//	var resps *HttpResponse
}

func readchannels() {
	xmlFile, err := os.Open("/etc/httpsServer/channel-list.xml")
	if err != nil {
		fmt.Println("Opening file error : ", err)
	}
	defer xmlFile.Close()
	xmlData, _ := ioutil.ReadAll(xmlFile)
	var A apps
	xml.Unmarshal(xmlData, &A)
	for _, value := range A.App {
		fmt.Printf("%s\n", value.App)
	}

}
