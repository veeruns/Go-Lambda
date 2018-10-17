package rokulib

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
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
var RokulibLog *os.File

//ChannelHash is the channel name to channel id map
var ChannelHash map[string]int
var Ljack lumberjack.Logger
var Log logrus.Logger

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
	var err error
	RokulibLog, err = os.OpenFile("/opt/httpsServer/logs/access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		Log.Fatal(err)
	}

	defer RokulibLog.Close()
	Ljack = lumberjack.Logger{
		Filename:   "/opt/httpsServer/logs/access.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default

	}
	Log.SetOutput(&Ljack)

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
	flag = false
	signal.Notify(signalchannel, syscall.SIGHUP)
	for {
		Log.Info("[Rokulib] Started reading from data channel")
		select {
		case msg := <-flagsignal:
			if flag {
				var resp *http.Response
				Log.Infof("[Rokulib]  Recieved to post %s\n ", msg)
				var buff bytes.Buffer
				resp, err := http.Post(msg, "", &buff)
				//	time.Sleep(time.Millisecond * 2500)
				if err != nil {
					Log.Infof("[Rokulib] Error from Roku %s\n", err.Error())
				} else {
					Log.Infof("[Rokulib] Response code from Roku %d\n", resp.StatusCode)
				}
			} else {
				Log.Infof("[Rokulib]  Recieved to post %s\n ", msg)
				time.Sleep(time.Millisecond * 2500)
				Log.Infof("Returning\n")
			}
		case <-time.After(30 * time.Second):
			Log.Infof("[Rokulib] Nothing recieved yet")
		case <-signalchannel:
			Ljack.Rotate()
			Log.Info("The log was rotated")
		}
	}
	//	var resps *HttpResponse
}

func readchannels() {
	xmlFile, err := os.Open("/etc/httpsServer/channel-list.xml")
	if err != nil {
		Log.Fatalf("Opening file error : %s", err.Error())
	}
	defer xmlFile.Close()
	xmlData, _ := ioutil.ReadAll(xmlFile)
	var A apps
	xml.Unmarshal(xmlData, &A)
	for _, value := range A.App {
		Log.Infof("%s\n", value.App)
		ChannelHash[value.App] = value.ID
	}

}
