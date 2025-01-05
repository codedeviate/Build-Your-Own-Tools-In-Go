package main

import (
	"fmt"
	"log"
	"os"

	"github.com/username/dhcpclient/lib"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: dhcpclient <network_interface>")
		os.Exit(1)
	}

	iface := os.Args[1]
	packet, err := lib.RequestIP(iface)
	if err != nil {
		log.Fatalf("Failed to request IP: %v", err)
	}

	fmt.Printf("Received DHCP offer: %v\n", packet.YourIPAddr)
}
