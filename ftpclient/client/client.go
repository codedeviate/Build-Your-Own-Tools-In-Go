// client.go
package client

import (
	"io"

	"github.com/jlaffaye/ftp"
)

// FetchFile fetches the content of the specified file from the FTP server.
func FetchFile(server, user, password, filePath string) (string, error) {
	conn, err := ftp.Dial(server)
	if err != nil {
		return "", err
	}
	defer conn.Quit()

	if err := conn.Login(user, password); err != nil {
		return "", err
	}

	resp, err := conn.Retr(filePath)
	if err != nil {
		return "", err
	}
	defer resp.Close()

	body, err := io.ReadAll(resp)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
