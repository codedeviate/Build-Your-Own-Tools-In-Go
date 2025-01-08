package main

import (
	"fmt"
	"log"
	"os"

	"github.com/username/certinfo/certificate"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <hostname>", os.Args[0])
	}

	hostname := os.Args[1]
	certInfo, err := certificate.FetchCertificateInfo(hostname)
	if err != nil {
		log.Fatalf("Error fetching certificate: %v", err)
	}

	fmt.Printf("Certificate Information for %s:\n", hostname)
	fmt.Printf("Issuer: %s\n", certInfo.Issuer)
	fmt.Printf("Expiry Date: %s\n", certInfo.ExpiryDate)
	fmt.Println("Certificate Chain:")
	for i, chainCert := range certInfo.Chain {
		fmt.Printf("  [%d] %s\n", i+1, chainCert.Subject)
	}
}
