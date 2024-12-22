// main.go
package main

import (
	"fmt"
	"os"
	"strconv"

	netscan "github.com/username/netscan/netscanmod"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: netscan <hostname> <startPort> <endPort>")
		os.Exit(1)
	}

	hostname := os.Args[1]
	startPort, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid start port: %v\n", err)
		os.Exit(1)
	}

	endPort, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Invalid end port: %v\n", err)
		os.Exit(1)
	}

	netscan.ScanHost(hostname, startPort, endPort)
}
