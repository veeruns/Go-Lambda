package rokulib

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
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
var signal chan string
var ChannelHash map[string]int

//ch := make(chan *HttpResponse, 1)

//InitLib initializes channel
func InitLib() {
	//	datachan := make(chan *HttpResponse, 1)
	signal = make(chan string)
	ChannelHash = make(map[string]int)
	readchannels()
	//start workerpool
	go workerpool()
	rokuliblog, err := os.OpenFile("/opt/httpsServer/logs/rokulib.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer rokuliblog.Close()

	log.SetOutput(&lumberjack.Logger{
		Filename:   "/opt/httpsServer/logs/rokulib.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

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

//LaunchChannel send the launch command to roku
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
	flag = false
	for {
		log.Info("[Rokulib] Started reading from data channel")
		select {
		case msg := <-signal:
			if flag {
				var resp *http.Response
				log.Infof("[Rokulib]  Recieved to post %s\n ", msg)
				var buff bytes.Buffer
				resp, err := http.Post(msg, "", &buff)
				//	time.Sleep(time.Millisecond * 2500)
				if err != nil {
					log.Infof("[Rokulib] Error from Roku %s\n", err.Error())
				} else {
					log.Infof("[Rokulib] Response code from Roku %d\n", resp.StatusCode)
				}
			} else {
				log.Infof("[Rokulib]  Recieved to post %s\n ", msg)
				time.Sleep(time.Millisecond * 2500)
				log.Infof("Returning\n")
			}
		case <-time.After(30 * time.Second):
			log.Infof("[Rokulib] Nothing recieved yet")

		}
	}
	//	var resps *HttpResponse
}

func readchannels() {
	xmlFile, err := os.Open("/etc/httpsServer/channel-list.xml")
	if err != nil {
		log.Fatalf("Opening file error : %s", err.Error())
	}
	defer xmlFile.Close()
	xmlData, _ := ioutil.ReadAll(xmlFile)
	var A apps
	xml.Unmarshal(xmlData, &A)
	for _, value := range A.App {
		log.Infof("%s\n", value.App)
		ChannelHash[value.App] = value.ID
	}

}
