// main.go
package main

import (
	"fmt"
	"os"
	"strconv"

	portscanner "github.com/username/portscanner/scanner"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: portscanner <protocol> <hostname> <startPort> <endPort>")
		os.Exit(1)
	}

	protocol := os.Args[1]
	hostname := os.Args[2]
	startPort, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Invalid start port: %v\n", err)
		os.Exit(1)
	}
	endPort, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Printf("Invalid end port: %v\n", err)
		os.Exit(1)
	}

	openPorts := portscanner.ScanPorts(protocol, hostname, startPort, endPort)
	if len(openPorts) == 0 {
		fmt.Println("No open ports found.")
	} else {
		fmt.Printf("Open ports: %v\n", openPorts)
	}
}
