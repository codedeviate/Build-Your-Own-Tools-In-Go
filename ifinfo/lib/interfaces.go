package lib

import (
	"fmt"
	"net"
)

// ListInterfaces retrieves all network interfaces and their IP addresses.
func ListInterfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Error retrieving interfaces: %v\n", err)
		return
	}

	fmt.Println("Available Network Interfaces:")
	for _, iface := range interfaces {
		fmt.Printf("Name: %s, MTU: %d, HardwareAddr: %s, Flags: %s\n", iface.Name, iface.MTU, iface.HardwareAddr, iface.Flags)

		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Printf("  Error retrieving addresses: %v\n", err)
			continue
		}

		for _, addr := range addrs {
			fmt.Printf("  Address: %s\n", addr.String())
		}
		fmt.Println()
	}
}
