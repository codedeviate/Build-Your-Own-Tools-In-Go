// client.go
package client

import (
	"crypto/tls"
	"io"

	"github.com/jlaffaye/ftp"
)

// FetchFile fetches the content of the specified file from the FTPS server.
func FetchFile(server, user, password, filePath string) (string, error) {
    // Configure TLS settings
    tlsConfig := &tls.Config{
        InsecureSkipVerify: true, // Skip certificate verification for simplicity
    }

    // Connect to the FTPS server
    conn, err := ftp.Dial(server, ftp.DialWithTLS(tlsConfig))
    if err != nil {
        return "", err
    }
    defer conn.Quit()

    // Login to the server
    if err := conn.Login(user, password); err != nil {
        return "", err
    }

    // Retrieve the file
    resp, err := conn.Retr(filePath)
    if err != nil {
        return "", err
    }
    defer resp.Close()

    // Read the file content
    body, err := io.ReadAll(resp)
    if err != nil {
        return "", err
    }

    return string(body), nil
}
