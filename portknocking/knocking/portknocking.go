// portknocking.go
package portknocking

import (
	"fmt"
	"net"
	"time"
)

// Knock sends a sequence of packets to the specified ports.
func Knock(host string, ports []int, delay time.Duration) error {
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.Dial("udp", address)
		if err != nil {
			return err
		}
		conn.Close()
		time.Sleep(delay)
	}
	return nil
}
