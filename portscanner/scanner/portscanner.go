// scanner.go
package portscanner

import (
	"fmt"
	"net"
	"time"
)

// ScanPort checks if a port is open on a given host.
func ScanPort(protocol, hostname string, port int) bool {
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol, address, 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// ScanPorts scans a range of ports on a given host.
func ScanPorts(protocol, hostname string, startPort, endPort int) []int {
	var openPorts []int
	for port := startPort; port <= endPort; port++ {
		if ScanPort(protocol, hostname, port) {
			openPorts = append(openPorts, port)
		}
	}
	return openPorts
}
