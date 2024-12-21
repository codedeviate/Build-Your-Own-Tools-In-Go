package main

import (
	"fmt"
	"github.com/username/packetsniffer/sniffer"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: packetsniffer <network_interface>")
		os.Exit(1)
	}

	device := os.Args[1]
	snapshotLen := int32(1024)
	promiscuous := false
	timeout := 30 * time.Second

	fmt.Printf("Starting packet sniffer on interface %s...\n", device)
	if err := sniffer.SniffPackets(device, snapshotLen, promiscuous, timeout); err != nil {
		log.Fatalf("Error sniffing packets: %v", err)
	}
}
