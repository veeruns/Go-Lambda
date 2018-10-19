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

	"github.com/spf13/viper"
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

type Config struct {
	Server        string
	Listenport    string
	Rokuurl       string
	Accesslogpath string
	Devflag       bool
	AuthString    string
	SSLcertname   string
	SSLkeyname    string
}

var Conf Config

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
	readconfig(&Conf)
	//start workerpool
	go workerpool()

}

//Readconfig File
func readconfig(*Config) bool {
	viper.SetConfigName("Server")
	viper.AddConfigPath("/opt/httpsServer/conf")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Config file not found...%s\n", err.Error())
		return false
	} else {
		//Server section
		Conf.Accesslogpath = viper.GetString("Server.accesslogpath")
		Conf.Server = viper.GetString("Server.Host")
		Conf.Listenport = viper.GetString("Server.ListenPort")
		//Roku section
		Conf.Devflag = viper.GetBool("Roku.Development")
		Conf.Rokuurl = viper.GetString("Roku.URL")
		//Authorization section
		Conf.AuthString = viper.GetString("Authorization.CommonName")
		//SSL Cert/keysection

		Conf.SSLcertname = viper.GetString("SSL.Certname")
		Conf.SSLkeyname = viper.GetString("SSL.KeyName")
	}
	return true
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
	flag = Conf.Devflag
	signal.Notify(signalchannel, syscall.SIGHUP)
	for {

		select {
		case msg := <-flagsignal:
			if flag {
				//var resp *http.Response

				var buff bytes.Buffer
				http.Post(msg, "", &buff)

			} else {

				time.Sleep(time.Millisecond * 2500)
				fmt.Printf("")
				fmt.Printf("Roku dev flag is %v so not doing anything\n", Conf.Devflag)

			}
		case <-time.After(30 * time.Second):
			//			fmt.Printf("Something ")
			fmt.Printf("Roku configuration %s\n", Conf.Rokuurl)
			fmt.Printf("Roku dev flag is %v\n", Conf.Devflag)

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
