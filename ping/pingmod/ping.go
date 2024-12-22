// ping.go
package ping

import (
	"net"
	"time"
)

// Ping sends a ping request to the specified address and returns the duration and any error encountered.
func Ping(address string) (time.Duration, error) {
    start := time.Now()
    conn, err := net.DialTimeout("ip4:icmp", address, time.Second*2)
    if err != nil {
        return 0, err
    }
    defer conn.Close()
    duration := time.Since(start)
    return duration, nil
}
