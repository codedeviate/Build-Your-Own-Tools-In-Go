package traceroute

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// Traceroute performs a traceroute to the specified address.
func Traceroute(address string, maxHops int) ([]string, error) {
	var hops []string

	targetAddr, err := net.ResolveIPAddr("ip4", address)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve address: %v", err)
	}

	for ttl := 1; ttl <= maxHops; ttl++ {
		rawConn, err := net.Dial("ip4:icmp", targetAddr.String())
		if err != nil {
			return nil, fmt.Errorf("failed to create connection: %v", err)
		}
		defer rawConn.Close()

		conn := ipv4.NewConn(rawConn)

		// Set TTL for this hop
		if err := conn.SetTTL(ttl); err != nil {
			return nil, fmt.Errorf("failed to set TTL: %v", err)
		}

		// Send ICMP Echo Request
		icmpMessage := icmp.Message{
			Type: ipv4.ICMPTypeEcho, Code: 0, Body: &icmp.Echo{
				ID:   ttl, // Use TTL as ID to simplify tracking
				Seq:  1,
				Data: []byte("traceroute"),
			},
		}

		icmpBytes, err := icmpMessage.Marshal(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal ICMP message: %v", err)
		}

		start := time.Now()
		if _, err := rawConn.Write(icmpBytes); err != nil {
			return nil, fmt.Errorf("failed to send ICMP message: %v", err)
		}

		// Wait for reply
		buffer := make([]byte, 1500)
		rawConn.SetReadDeadline(time.Now().Add(2 * time.Second))
		n, err := rawConn.Read(buffer)
		peer := rawConn.RemoteAddr()

		checkTTL, _ := conn.TTL()
		fmt.Printf("%d, %d, %v, %v\n", ttl, checkTTL, rawConn.RemoteAddr().String(), buffer[:])

		if err != nil {
			// Append timeout if no response
			hops = append(hops, fmt.Sprintf("Hop %d: Request timed out (%s)", ttl, err.Error()))
			continue
		}

		duration := time.Since(start)
		reply, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), buffer[:n])
		if err != nil {
			return nil, fmt.Errorf("failed to parse ICMP reply: %v", err)
		}

		if reply.Type == ipv4.ICMPTypeEchoReply {
			hops = append(hops, fmt.Sprintf("Hop %d: %v (%v) in %v", ttl, peer, targetAddr, duration))
			if peer.String() == targetAddr.String() {
				break // Reached destination
			}
		} else {
			hops = append(hops, fmt.Sprintf("Hop %d: Unexpected ICMP type %v", ttl, reply.Type))
			if reply.Type == ipv4.ICMPTypeDestinationUnreachable {
				break // Reached destination
			}
		}
	}

	return hops, nil
}
