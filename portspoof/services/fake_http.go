// services/fake_http.go
package services

import (
	"crypto/tls"
)

// FakeHTTPService simulates an HTTP service.
type FakeHTTPService struct {
	*GenericService
}

// NewFakeHTTPService returns a new FakeHTTPService.
func NewFakeHTTPService(customBanner string, useTLS bool, tlsConfig *tls.Config) *FakeHTTPService {
	banner := customBanner
	if banner == "" {
		banner = "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n<html><body><h1>Fake HTTP Service</h1></body></html>"
	}
	return &FakeHTTPService{
		GenericService: NewGenericService("http", banner, useTLS, tlsConfig),
	}
}
