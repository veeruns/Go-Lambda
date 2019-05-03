package main

import (
	"flag"
	"log"
	"os"
	"git.ouroath.com/hga/athenz-user-cert"
)

func main() {
	pIdentity := flag.String("i", os.Getenv("USER"), "the identity to authenticate as")
	pReset := flag.Bool("r", false, "reset existing key/cert files")
	pServer := flag.String("s", "", "sscha athenz proxy server name")
	pPort := flag.Int("p", 10022, "sshca athenz proxy server port")

	// first we need to parse our arguments based
	// on the flags we defined above

	flag.Parse()

	if *pIdentity == "" || *pServer == "" {
		log.Fatalf("usage: athenz-user-cert -i <user> -s <server> -p <port>")
	}

	// if we're ask to reset first we'll remove the
	// existing key and certificate files

	if *pReset {
		athenzusercert.ResetUser()
	}

	// request new user cert from ZTS

	_, _, err := athenzusercert.RetrieveUserCert(*pIdentity, *pServer, *pPort)
	if err != nil {
		log.Fatalf("Failed to get certificate: %v\nMake sure you have run 'yinit --hard' before requesting a user certificate", err)
	}
	log.Println("Successfully retrieved user Athenz x.509 certificate")
}
