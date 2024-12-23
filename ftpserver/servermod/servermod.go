// server.go
package servermod

import (
	"strconv"

	"goftp.io/server/v2"
	"goftp.io/server/v2/driver/file"
)

// StartServer starts the FTP server on the specified port.
func StartServer(port string, path string) error {
	driver, err := file.NewDriver(path)

	if err != nil {
		return err
	}

	portValue, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	s, err := server.NewServer(&server.Options{
		Driver: driver,
		Auth: &server.SimpleAuth{
			Name:     "admin",
			Password: "admin",
		},
		Perm:      server.NewSimplePerm("root", "root"),
		RateLimit: 1000000, // 1MB/s limit
		Port:      portValue,
	})
	if err != nil {
		return err
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
