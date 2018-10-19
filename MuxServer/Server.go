package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	//	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"github.com/veeruns/Go-Lambda/rokulib"
)

type config struct {
	server        string
	listenport    string
	rokuurl       string
	accesslogpath string
	devflag       bool
}

func RokuServer(w http.ResponseWriter, req *http.Request) {
	var functocall string
	var channeltocall string
	functocall = req.URL.Query().Get("func")
	channeltocall = req.URL.Query().Get("channel")
	log.Infof("Raw url string %s\n", req.URL.RequestURI())
	wholeurl, _ := url.Parse(req.URL.RequestURI())
	//queryparams := wholeurl.Query()
	log.Infof("The whole url is %s\n", wholeurl.String())
	log.Infof("The query function is %s\n", functocall)
	log.Warnf("Channel change is %s\n", channeltocall)
	m, _ := url.ParseQuery(wholeurl.RawQuery)
	log.Infof("Parameters are %s\n", m)
	log.Infof("Whole query is %s\n", m["channel"])
	if len(req.TLS.PeerCertificates) > 0 {
		fmt.Fprintf(w, "client common name: %+v\n", req.TLS.PeerCertificates[0].Subject.CommonName)
		fmt.Fprintf(w, "Client OU %+v\n", req.TLS.PeerCertificates[0].Subject.OrganizationalUnit)
		log.Infof("client common name: %+v\n", req.TLS.PeerCertificates[0].Subject.CommonName)
		log.Infof("Client OU %+v\n", req.TLS.PeerCertificates[0].Subject.OrganizationalUnit)
		//	fmt.Fprintf(w, " %s\n", req.TLS.PeerCertificates[0].Verify)
	}
	for _, certname := range req.TLS.PeerCertificates {
		for _, dnsname := range certname.DNSNames {
			log.Infof("SAN names listed %s\n", dnsname)
		}
	}
	if strings.Compare(req.TLS.PeerCertificates[0].Subject.CommonName, rokulib.Conf.AuthString) != 0 {

		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "meh")
	} else {
		w.Header().Set("Content-Type", "text/plain")
		var works bool
		switch {
		case strings.Compare(functocall, "off") == 0:
			works = rokulib.PowerOff("192.168.7.45:8060")
			log.Infof("Rokulib returned %v\n", works)
		case strings.TrimRight(functocall, "\n") == "on":
			works = rokulib.PowerOn("192.168.7.45:8060")
			log.Infof(" Rokulib PowerOn returned %v\n", works)
		case len(channeltocall) > 0:
			channelname := strings.TrimRight(channeltocall, "\n")
			valuid := rokulib.ChannelHash[channelname]
			log.Infof("Channel Name is %s and Channel id is %d\n", channelname, valuid)
			works = rokulib.LaunchChannel("192.168.7.45:8060", valuid)
			log.Infof("Rokulib returned the %v\n", works)
		default:
			works = false
			log.Infof("We are calling default\n")
		}

		//"http://192.168.7.45:8060/keypress/powerOff",
		var retbytes bytes.Buffer
		retbytes.WriteString(functocall)
		if works {
			log.Infof("It Returned a bool")
			retbytes.WriteString(" Achieved")
		}

		io.WriteString(w, retbytes.String())

		//rokulib.GetResponses()
	}
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	rokulib.InitLib()
	var sigchannel chan os.Signal
	sigchannel = make(chan os.Signal, 1)
	accesslog, err := os.OpenFile(rokulib.Conf.Log, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
		log.Fatal(err.Error())
	}
	defer accesslog.Close()

	var ljack lumberjack.Logger
	ljack = lumberjack.Logger{
		Filename:   rokulib.Conf.Log,
		MaxSize:    5, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
	mWriter := io.MultiWriter(accesslog, ljack)
	log.SetOutput(mWriter)
	signal.Notify(sigchannel, syscall.SIGHUP)
	go func() {
		for {
			<-sigchannel
			ljack.Rotate()
			log.Info("accesslog is rotated")
		}
	}()

	mux := mux.NewRouter()

	log.Infof("Server Configuration %s", rokulib.Conf.Server)
	log.Infof("Server port %s", rokulib.Conf.Listenport)
	log.Infof("Server access log path %s", rokulib.Conf.Log)
	mux.HandleFunc("/roku", RokuServer)
	loggedRouter := handlers.CombinedLoggingHandler(mWriter, mux)
	//mux.Use(handlers.CombinedLoggingHandler(os.StdOut, ))
	caCert, err := ioutil.ReadFile(rokulib.Conf.CAcert)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		MinVersion:         tls.VersionTLS12,
		ClientAuth:         tls.RequireAndVerifyClientCert,
		InsecureSkipVerify: true,
		ClientCAs:          caCertPool,
	}
	var v bytes.Buffer
	v.WriteString(rokulib.Conf.Server)
	v.WriteString(":")
	v.WriteString(rokulib.Conf.Listenport)
	srv := &http.Server{
		Addr: v.String(),
		//		Handler:   mux,
		TLSConfig: cfg,
		Handler:   loggedRouter,
	}

	srv.ListenAndServeTLS(rokulib.Conf.SSLcertname, rokulib.Conf.SSLkeyname)

}

/*
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		tag := fmt.Sprintf("[%s -> %s]", req.URL, req.RemoteAddr)
		log.Infof(format)("%s accept", tag)

		if len(req.TLS.PeerCertificates) > 0 {
			log.Infof("%s client common name: %+v", tag, req.TLS.PeerCertificates[0].Subject.CommonName)
		}

		next.ServeHTTP(w, req)
	})
}*/
