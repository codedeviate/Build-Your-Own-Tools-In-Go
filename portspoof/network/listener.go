// network/listener.go
package network

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"portspoof/config"
	"portspoof/services"
	"strconv"
)

// StartListeners creates a listener for every service defined in the config.
func StartListeners(cfg config.Config) {
	// Load shared TLS configuration if available.
	tlsConfig, err := loadTLSConfig(cfg.TLSCertFile, cfg.TLSKeyFile)
	if err != nil {
		log.Printf("TLS configuration load error: %v. TLS will be disabled.", err)
		tlsConfig = nil
	}

	for _, ps := range cfg.Services {
		// Randomly decide whether to use TLS (if available) for this service.
		useTLS := (rand.Intn(2) == 0) && (tlsConfig != nil)

		var svc services.Service
		switch ps.Service {
		case "ssh":
			svc = services.NewFakeSSHService(ps.Banner, useTLS, tlsConfig)
		case "http":
			svc = services.NewFakeHTTPService(ps.Banner, useTLS, tlsConfig)
		case "mysql":
			svc = services.NewFakeMySQLService(ps.Banner, useTLS, tlsConfig)
		default:
			banner := ps.Banner
			if banner == "" {
				banner = fmt.Sprintf("Welcome to fake service on port %d", ps.Port)
			}
			svc = services.NewGenericService("generic", banner, useTLS, tlsConfig)
		}
		go startListener(ps.Port, svc, useTLS, tlsConfig)
	}
}

// startListener binds to the given port and services incoming connections.
func startListener(port int, svc services.Service, useTLS bool, tlsConfig *tls.Config) {
	addr := "[::]:" + strconv.Itoa(port)
	var ln net.Listener
	var err error

	if useTLS && tlsConfig != nil {
		ln, err = tls.Listen("tcp", addr, tlsConfig)
	} else {
		ln, err = net.Listen("tcp", addr)
	}
	if err != nil {
		log.Printf("Error listening on %s: %v", addr, err)
		return
	}
	log.Printf("Listening on %s for service %s (TLS: %v)", addr, svc.Banner(), useTLS)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Accept error on port %d: %v", port, err)
			continue
		}
		remoteIP, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
		go svc.Handle(conn, remoteIP)
	}
}

// loadTLSConfig attempts to load the TLS certificate and key.
func loadTLSConfig(certFile, keyFile string) (*tls.Config, error) {
	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("TLS certificate file %s not found", certFile)
	}
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}
