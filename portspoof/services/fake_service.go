// services/fake_service.go
package services

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"portspoof/alerts"
	"portspoof/dashboard"
	"portspoof/logging"
	"strings"
	"time"
)

// Service is the interface for a fake service.
type Service interface {
	Handle(conn net.Conn, remoteIP string)
	Banner() string
}

// GenericService implements common functionality for fake services.
type GenericService struct {
	serviceName string
	banner      string
	useTLS      bool
	tlsConfig   *tls.Config
}

// NewGenericService creates a new GenericService.
func NewGenericService(name, banner string, useTLS bool, tlsConfig *tls.Config) *GenericService {
	return &GenericService{
		serviceName: name,
		banner:      banner,
		useTLS:      useTLS,
		tlsConfig:   tlsConfig,
	}
}

// Banner returns the configured banner.
func (s *GenericService) Banner() string {
	return s.banner
}

// Handle processes a connection: sends a banner, logs input, and echoes responses.
func (s *GenericService) Handle(conn net.Conn, remoteIP string) {
	defer conn.Close()

	logging.LogEvent(s.serviceName, fmt.Sprintf("Connection from %s", remoteIP))
	dashboard.AddEvent(dashboard.DashboardEvent{
		Time:     time.Now(),
		Service:  s.serviceName,
		Message:  "New connection",
		RemoteIP: remoteIP,
	})

	if s.useTLS {
		if tlsConn, ok := conn.(*tls.Conn); ok {
			if err := tlsConn.Handshake(); err != nil {
				logging.LogEvent(s.serviceName, fmt.Sprintf("TLS handshake error from %s: %v", remoteIP, err))
				return
			}
		}
	}

	// Send the fake banner.
	conn.Write([]byte(s.Banner() + "\n"))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		input := scanner.Text()
		logging.LogEvent(s.serviceName, fmt.Sprintf("Input from %s: %s", remoteIP, input))
		dashboard.AddEvent(dashboard.DashboardEvent{
			Time:     time.Now(),
			Service:  s.serviceName,
			Message:  "Received input: " + input,
			RemoteIP: remoteIP,
		})
		alerts.RecordActivity(remoteIP)
		// Echo back a simple response.
		response := strings.ToUpper(input)
		conn.Write([]byte("Echo: " + response + "\n"))
	}
}
