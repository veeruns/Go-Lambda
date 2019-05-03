package main

import (
	"crypto/tls"
	"fmt"

	"git.ouroath.com/go/ytls"
	ykk "git.ouroath.com/ykeykey/go/ykeykey"
)

func getTLSConfigFromFiles(certFile, keyFile string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	config := ytls.ClientTLSConfig()
	config.Certificates = []tls.Certificate{cert}

	// Set Renegotiation explicitly
	config.Renegotiation = tls.RenegotiateOnceAsClient

	return config, nil
}

func main() {
	tlsConfig, err := getTLSConfigFromFiles("/Users/veeru/.athenz/cert", "/Users/veeru/.athenz/key")
	if err != nil {
		return nil, fmt.Errorf("Unable to formulate tls config, error: %v", err)
	}

	transportOpts := ykk.TransportOptions{
		TLSClientConfig: tlsConfig,
	}

	// Auth Providers can be nil since we are using TLS certificates.
	// Initialize the slice with your ykk.AuthProvider implementations
	// if needed.
	var authProviders []ykk.AuthProvider

	client, err := ykk.NewSecretClient(
		ykk.Environment(env),
		ykk.Options{Transport: transportOpts},
		authProviders,
	)
	group := "commsensemble.team.secrets"
	key := "snow"
	value, err := client.Key(group, key)
	fmt.Printf("Key is %s\n", value)
}
