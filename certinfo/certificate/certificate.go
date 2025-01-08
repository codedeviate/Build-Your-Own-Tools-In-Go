package certificate

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"time"
)

// CertInfo holds key certificate details.
type CertInfo struct {
	Issuer     string
	ExpiryDate string
	Chain      []*x509.Certificate
}

// FetchCertificateInfo connects to a host and extracts SSL certificate details.
func FetchCertificateInfo(hostname string) (*CertInfo, error) {
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", hostname), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", hostname, err)
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return nil, fmt.Errorf("no certificates found for %s", hostname)
	}

	// Extract main certificate details
	mainCert := certs[0]
	issuer := mainCert.Issuer.String()
	expiryDate := mainCert.NotAfter.Format(time.RFC1123)

	// Return all certificates in the chain
	return &CertInfo{
		Issuer:     issuer,
		ExpiryDate: expiryDate,
		Chain:      certs,
	}, nil
}
