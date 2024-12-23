// client.go
package client

import (
	"io"
	"net/http"
)

// FetchURL fetches the content of the specified URL and returns it as a string.
func FetchURL(url string) (string, error) {
	resp, err := http.Get(url)
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
