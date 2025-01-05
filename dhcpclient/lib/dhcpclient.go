// dhcpclient.go
package lib

import (
	"context"
	"time"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/nclient4"
)

// RequestIP requests an IP address from a DHCP server.
func RequestIP(iface string) (*dhcpv4.DHCPv4, error) {
    client, err := nclient4.New(iface, nclient4.WithTimeout(10*time.Second))
    if err != nil {
        return nil, err
    }
    defer client.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    packet, err := client.DiscoverOffer(ctx)
    if err != nil {
        return nil, err
    }

    return packet, nil
}
