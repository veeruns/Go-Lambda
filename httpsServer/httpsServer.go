package main

import (
	// "fmt"
	// "io"

	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	addr string
}

// NewServer return a new echo server
func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func RokuServer(w http.ResponseWriter, req *http.Request) {

	if len(req.TLS.PeerCertificates) > 0 {
		fmt.Fprintf(w, "client common name: %+v\n", req.TLS.PeerCertificates[0].Subject.CommonName)
		fmt.Fprintf(w, "Client OU %+v\n", req.TLS.PeerCertificates[0].Subject.OrganizationalUnit)
		fmt.Fprintf(w, " %+v\n", req.TLS.PeerCertificates[0].Verify)
	}
	for _, certname := range req.TLS.PeerCertificates {
		for _, dnsname := range certname.DNSNames {
			fmt.Printf("SAN names listed %s\n", dnsname)
		}
	}
	if strings.Compare(req.TLS.PeerCertificates[0].Subject.CommonName, "client-auth.raghavanonline.com") != 0 {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "You do not have access to this server")
	} else {
		w.Header().Set("Content-Type", "text/plain")

		var works bool
		works = PowerOn("192.168.7.45:8060")

		//"http://192.168.7.45:8060/keypress/powerOff",
		if err2 != nil {
			fmt.Printf("Did not work as expected %s\n", err2.Error())
		} else {
			fmt.Printf("%s\n", resp.Body)
		}
		w.Write([]byte("This is an example server.\n"))
	}
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	http.HandleFunc("/roku", RokuServer)
	caCert, err := ioutil.ReadFile("ssl/certs/CAcerts.pem")
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

func PowerOff(hostname string) bool {
	var buff bytes.Buffer
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/keypress/poweron")
	resp, err := http.Post(url.String(), "", &buff)
	if err != nil {
		fmt.Printf("That did not work as intended %s\n", err.Error())
		return false
	} else {
		fmt.Printf("The return string is %s\n", resp.Body)
		return true
	}
}

func PowerOn(hostname string) bool {
	var buff bytes.Buffer
	var url bytes.Buffer
	url.WriteString("http://")
	url.WriteString(hostname)
	url.WriteString("/keypress/poweron")
	resp, err := http.Post(url.String(), "", &buff)
	if err != nil {
		fmt.Printf("That did not work as intended %s\n", err.Error())
		return false
	} else {
		fmt.Printf("The return string is %s\n", resp.Body)
		return true
	}
}
