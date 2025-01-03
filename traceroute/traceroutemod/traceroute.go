package traceroute

import (
	"fmt"
	"net"
	"strings"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// Config struct holds configuration values for the traceroute operation
// including whether to resolve DNS names and the maximum number of hops and timeout.
// The struct also includes flags to skip hops 1-5 which normally can cause timeouts.
type Config struct {
	Dns      *bool    // DNS resolution flag for hop addresses
	Hops     *int     // Maximum number of hops
	Timeout  *float64 // Timeout in seconds for each hop
	SkipHop1 *bool    // Skip hop 1
	SkipHop2 *bool    // Skip hop 2
	SkipHop3 *bool    // Skip hop 3
	SkipHop4 *bool    // Skip hop 4
	SkipHop5 *bool    // Skip hop 5
}

// Traceroute performs a traceroute to the specified address.
// It sends ICMP Echo requests to the target IP and prints the route
// taken by the packets, showing the response times for each hop.
func Traceroute(address string, config Config) error {
	// Resolve the target address to an IP address
	targetAddr, err := net.ResolveIPAddr("ip4", address)
	if err != nil {
		return fmt.Errorf("failed to resolve address: %v", err)
	}

	// Loop over the TTL values to simulate the hop count
	for ttl := 1; ttl <= *config.Hops; ttl++ {
		if ttl == 1 && *config.SkipHop1 {
			fmt.Println("Hop 1: Skipped")
			continue
		} else if ttl == 2 && *config.SkipHop2 {
			fmt.Println("Hop 2: Skipped")
			continue
		} else if ttl == 3 && *config.SkipHop3 {
			fmt.Println("Hop 3: Skipped")
			continue
		} else if ttl == 4 && *config.SkipHop4 {
			fmt.Println("Hop 4: Skipped")
			continue
		} else if ttl == 5 && *config.SkipHop5 {
			fmt.Println("Hop 5: Skipped")
			continue
		}

		// Listen for ICMP packets on the network
		conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
		if err != nil {
			return fmt.Errorf("failed to create connection: %v", err)
		}
		defer conn.Close() // Ensure the connection is closed when done

		// Create an IPv4 packet connection for sending and receiving
		ipv4conn := conn.IPv4PacketConn()
		defer ipv4conn.Close() // Ensure IPv4 packet connection is closed when done

		// Set the TTL value for the outgoing packet
		if err := ipv4conn.SetTTL(ttl); err != nil {
			return fmt.Errorf("failed to set TTL: %v", err)
		}
		ipv4conn.SetControlMessage(ipv4.FlagTTL, true) // Enable TTL information in the control message

		// Create the ICMP Echo Request message
		message := []byte("traceroute") // Data for the ICMP packet
		icmpMessage := icmp.Message{
			Type: ipv4.ICMPTypeEcho, // Set the message type as Echo Request
			Code: 0,
			Body: &icmp.Echo{
				ID:   ttl, // Use TTL as ID to simplify tracking
				Seq:  1,
				Data: message,
			},
		}

		// Marshal the ICMP message into a byte slice
		icmpBytes, err := icmpMessage.Marshal(nil)
		if err != nil {
			return fmt.Errorf("failed to marshal ICMP message: %v", err)
		}

		// Record the time before sending the request
		start := time.Now()

		// Send the ICMP message to the target IP address
		if _, err := ipv4conn.WriteTo(icmpBytes, nil, &net.UDPAddr{IP: targetAddr.IP}); err != nil {
			return fmt.Errorf("failed to send ICMP message: %v", err)
		}

		// Wait for a reply from the target
		buffer := make([]byte, 1500)

		// Set the timeout duration in milliseconds and apply it to the connection read deadline
		timeout := int(*config.Timeout*1000) * int(time.Millisecond)
		err = conn.SetReadDeadline(time.Now().Add(time.Duration(timeout)))

		if err != nil {
			return fmt.Errorf("failed to set read deadline: %v", err)
		}

		// Read the reply from the target (or timeout if no response)
		n, cm, _, err := ipv4conn.ReadFrom(buffer)
		if err != nil {
			// If no response within the timeout, print a timeout message
			fmt.Printf("Hop %d: Request timed out\n", ttl)
			continue
		}

		// Calculate the round-trip duration
		duration := time.Since(start)

		// Parse the received ICMP Echo Reply message
		_, err = icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), buffer[:n])
		if err != nil {
			return fmt.Errorf("failed to parse ICMP reply: %v", err)
		}

		// Extract the source address and return TTL from the control message
		remoteAddress := "?.?.?.?" // Default value if no control message
		returnTTL := -1
		if cm != nil {
			remoteAddress = cm.Src.String() // Source IP address of the reply
			returnTTL = cm.TTL              // TTL returned by the source
		}
		orgRemoteAddress := remoteAddress // Store the original remote address

		// Optionally resolve the DNS name for the remote address if the flag is set
		if *config.Dns {
			names, err := net.LookupAddr(remoteAddress)
			if err == nil && len(names) > 0 {
				remoteAddress = fmt.Sprintf("%s (%s)", strings.Trim(names[0], " ."), remoteAddress)
			}
		}

		// Print the hop information, including the remote address and round-trip time
		fmt.Printf("Hop %d: %v in %v (Return TTL: %d)\n", ttl, remoteAddress, duration, returnTTL)

		// Close the connection before the next iteration
		// By closing the connection, we ensure that the next iteration will use a new socket
		// and the time to close the connection isn't postponed until the end of the function
		// which could cause a delay before the function returns. By using a go routine, we can
		// close the connection asynchronously while the next iteration is running.
		go ipv4conn.Close()

		// If the response address matches the target, we have reached the destination
		if orgRemoteAddress == targetAddr.String() {
			break // Reached destination
		}
	}

	return nil
}
