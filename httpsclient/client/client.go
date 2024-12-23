// client.go
package client

import (
	"crypto/tls"
	"io"
	"net/http"
)

// FetchURL fetches the content of the specified URL and returns it as a string.
func FetchURL(url string) (string, error) {
	// Create a custom HTTP client with TLS configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Skip certificate verification for simplicity
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
