// dnsclient.go
package client

import (
	"context"
	"net"
	"time"
)

// ResolveDNS resolves the given hostname to its IP addresses.
func ResolveDNS(hostname string, serverAddress string) ([]net.IP, error) {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, serverAddress)
		},
	}

	ips, err := r.LookupIP(context.Background(), "ip", hostname)
	if err != nil {
		return nil, err
	}
	return ips, nil
}
