// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/username/packetsniffer/proxy"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: proxyserver <listen_port> <target>")
		os.Exit(1)
	}

	listenPort := os.Args[1]
	target := os.Args[2]

	fmt.Printf("Starting proxy server on port %s, forwarding to %s...\n", listenPort, target)
	if err := proxy.StartProxy(listenPort, target); err != nil {
		log.Fatalf("Proxy server failed: %v", err)
	}
}
