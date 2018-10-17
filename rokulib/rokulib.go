package rokulib

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

//HTTPResponse a more complex type
type HTTPResponse struct {
	url  string
	resp *http.Response
	err  error
}

var flag bool

type app struct {
	ID          int    `xml:"id,attr"`
	Version     string `xml:"version,attr"`
	Channeltype string `xml:"type,attr"`
	App         string `xml:",chardata"`
}

type apps struct {
	XMLName xml.Name `xml:"apps"`
	App     []app    `xml:"app"`
}

//var datachan chan *HttpResponse
var flagsignal chan string
var signalchannel chan os.Signal

//RokulibLog file for server
var RokulibLog *os.File

//ChannelHash is the channel name to channel id map
var ChannelHash map[string]int

//ch := make(chan *HttpResponse, 1)

//InitLib initializes channel
func InitLib() {
	//	datachan := make(chan *HttpResponse, 1)
	flagsignal = make(chan string)
	signalchannel = make(chan os.Signal, 1)
	ChannelHash = make(map[string]int)
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
	flagsignal <- url.String()
	return true
}

//PowerOn function powers on roku
func PowerOn(hostname string) bool {
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/keypress/poweron")
	flagsignal <- url.String()
	return true
}

//LaunchChannel send the launch command to roku
func LaunchChannel(hostname string, channelid int) bool {
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/launch/")
	url.WriteString(strconv.Itoa(channelid))
	flagsignal <- url.String()
	return true

}

func workerpool() {
	flag = true
	signal.Notify(signalchannel, syscall.SIGHUP)
	for {

		select {
		case msg := <-flagsignal:
			if flag {
				var resp *http.Response

				var buff bytes.Buffer
				http.Post(msg, "", &buff)

			} else {

				time.Sleep(time.Millisecond * 2500)

			}
		case <-time.After(30 * time.Second):
			fmt.Printf("Something ")

		}
	}
	//	var resps *HttpResponse
}

func readchannels() {
	xmlFile, err := os.Open("/etc/httpsServer/channel-list.xml")
	if err != nil {
		//	Log.Fatalf("Opening file error : %s", err.Error())
	}
	defer xmlFile.Close()
	xmlData, _ := ioutil.ReadAll(xmlFile)
	var A apps
	xml.Unmarshal(xmlData, &A)
	for _, value := range A.App {
		//	Log.Infof("%s\n", value.App)
		ChannelHash[value.App] = value.ID
	}

}
