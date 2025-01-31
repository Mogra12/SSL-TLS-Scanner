package main

import (
	"crypto/tls"
	"flag"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	var hostname string

	vermelho := color.New(color.FgRed).SprintFunc()
	verde := color.New(color.FgGreen).SprintFunc()

	flag.StringVar(&hostname, "h", "", "Defines hostname")
	flag.Parse()

	if hostname == "" {
		fmt.Println("The hostname is invalid or not defined!")
	}

	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", hostname, config)
	if err != nil {
		fmt.Println(vermelho("ERROR"), ": Connection error, or invalid hostname. Hostname must be like this:", verde("<hostname:port>"))
		return
	}
	defer conn.Close()

	state := conn.ConnectionState()
	cert := state.PeerCertificates[0]

	fmt.Println("=== SSL Certificate ===")
	fmt.Println("Issued to:", cert.Subject.CommonName)
	fmt.Println("Issued by:", cert.Issuer.CommonName)
	fmt.Println("Valid from:", cert.NotBefore)
	fmt.Println("Valid until:", cert.NotAfter)
	fmt.Println("Signature algorithm:", cert.SignatureAlgorithm)

	fmt.Println("\n=== TLS Configuration ===")
	fmt.Println("TLS Version:", getTLSVersion(state.Version))
	fmt.Println("Cipher Suite:", tls.CipherSuiteName(state.CipherSuite))

}

func getTLSVersion(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown"
	}
}
