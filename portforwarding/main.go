// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/username/portforwarding/forwarder"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: portforwarding <listen_port> <target>")
		os.Exit(1)
	}

	listenPort := os.Args[1]
	target := os.Args[2]

	fmt.Printf("Starting port forwarder on port %s, forwarding to %s...\n", listenPort, target)
	if err := forwarder.StartForwarder(listenPort, target); err != nil {
		log.Fatalf("Port forwarder failed: %v", err)
	}
}
