package main

import (
	// "fmt"
	// "io"

	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/veeruns/Go-Lambda/rokulib"
)

type Server struct {
	addr string
}

// NewServer return a new echo server
func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

//RokuServer is the handler for https
func RokuServer(w http.ResponseWriter, req *http.Request) {
	var functocall string
	var channeltocall string
	functocall = req.URL.Query().Get("func")
	channeltocall = req.URL.Query().Get("channel")
	fmt.Printf("Raw url string %s\n", req.URL.RequestURI())
	wholeurl, _ := url.Parse(req.URL.RequestURI())
	//queryparams := wholeurl.Query()
	fmt.Printf("The whole url is %s\n", wholeurl.String())
	fmt.Printf("The query function is %s\n", functocall)
	fmt.Printf("Channel change is %s\n", channeltocall)
	m, _ := url.ParseQuery(wholeurl.RawQuery)
	fmt.Println("Parameters are ", m)
	fmt.Printf("Whole query is %s\n", m["channel"])
	if len(req.TLS.PeerCertificates) > 0 {
		fmt.Fprintf(w, "client common name: %+v\n", req.TLS.PeerCertificates[0].Subject.CommonName)
		fmt.Fprintf(w, "Client OU %+v\n", req.TLS.PeerCertificates[0].Subject.OrganizationalUnit)
		fmt.Printf("client common name: %+v\n", req.TLS.PeerCertificates[0].Subject.CommonName)
		fmt.Printf("Client OU %+v\n", req.TLS.PeerCertificates[0].Subject.OrganizationalUnit)
		//	fmt.Fprintf(w, " %s\n", req.TLS.PeerCertificates[0].Verify)
	}
	for _, certname := range req.TLS.PeerCertificates {
		for _, dnsname := range certname.DNSNames {
			fmt.Printf("SAN names listed %s\n", dnsname)
		}
	}
	if strings.Compare(req.TLS.PeerCertificates[0].Subject.CommonName, "client-auth.raghavanonline.com") != 0 {

		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "meh")
	} else {
		w.Header().Set("Content-Type", "text/plain")
		var works bool
		switch {
		case strings.Compare(functocall, "off") == 0:
			works = rokulib.PowerOff("192.168.7.45:8060")
			fmt.Printf(" Rokulib PowerOff returned %v\n", works)
		case strings.TrimRight(functocall, "\n") == "on":
			works = rokulib.PowerOn("192.168.7.45:8060")
			fmt.Printf(" Rokulib PowerOn returned %v\n", works)
		case len(channeltocall) > 0:
			channelname := strings.TrimRight(channeltocall, "\n")
			valuid := rokulib.ChannelHash[channelname]
			fmt.Printf("Channel Name is %s and Channel id is %d\n", channelname, valuid)
			works = rokulib.LaunchChannel("192.168.7.45:8060", valuid)
			fmt.Printf("Rokulib returned the %v\n", works)
		default:
			works = false
			fmt.Printf("We are calling default\n")
		}

		//"http://192.168.7.45:8060/keypress/powerOff",
		var retbytes bytes.Buffer
		retbytes.WriteString(functocall)
		if works {
			fmt.Printf("It Returned a bool")
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
	http.HandleFunc("/roku", RokuServer)

	caCert, err := ioutil.ReadFile("/etc/httpsServer/ssl/certs/CAcerts.pem")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{
		MinVersion:         tls.VersionTLS12,
		ClientAuth:         tls.RequireAndVerifyClientCert,
		InsecureSkipVerify: true,
		ClientCAs:          caCertPool,
	}
	srv := &http.Server{
		Addr: ":8080",
		//		Handler:   mux,
		TLSConfig: cfg,
	}

	log.Fatal(srv.ListenAndServeTLS("/etc/letsencrypt/live/veeruns.raghavanonline.com/cert.pem", "/etc/letsencrypt/live/veeruns.raghavanonline.com/privkey.pem"))
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		tag := fmt.Sprintf("[%s -> %s]", req.URL, req.RemoteAddr)
		log.Printf("%s accept", tag)

		if len(req.TLS.PeerCertificates) > 0 {
			log.Printf("%s client common name: %+v", tag, req.TLS.PeerCertificates[0].Subject.CommonName)
		}

		next.ServeHTTP(w, req)
	})
}
