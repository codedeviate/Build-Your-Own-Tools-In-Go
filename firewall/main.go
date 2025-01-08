// main.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/username/firewall/lib"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: firewall <network_interface>")
		os.Exit(1)
	}

	iface := os.Args[1]

	physicalIface, err := net.InterfaceByName(iface)
	if err != nil {
		log.Fatalf("could not find interface %s: %v", iface, err)
	}

	addrs, err := physicalIface.Addrs()
	if err != nil {
		log.Fatalf("could not get addresses for interface %s: %v", iface, err)
	}

	fw := lib.NewFirewall()

	for _, addr := range addrs {
		// Check if the address is an IP address
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		// Skip loopback addresses
		if ip == nil || ip.IsLoopback() {
			continue
		}

		fw.AddRule(lib.Rule{
			SrcIP:   ip,
			DstIP:   net.IPv4(8, 8, 8, 8),
			SrcPort: 53,
			DstPort: 53,
			Action:  "allow",
		})
	}

	fmt.Printf("Starting firewall on interface %s...\n", iface)
	lib.StartFirewall(iface, fw)
}
