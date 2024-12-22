// netscan.go
package netscan

import (
	"fmt"
	"net"
	"time"
)

// ScanPort checks if a port is open on a given host.
func ScanPort(protocol, hostname string, port int) bool {
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol, address, time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// ScanHost scans a range of ports on a given host.
func ScanHost(hostname string, startPort, endPort int) {
	for port := startPort; port <= endPort; port++ {
		if ScanPort("tcp", hostname, port) {
			fmt.Printf("Port %d is open\n", port)
		}
	}
}
