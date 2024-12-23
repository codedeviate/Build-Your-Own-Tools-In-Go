// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/dnsclient/client"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: dnsclient <hostname> [<server>]")
		os.Exit(1)
	}

	hostname := os.Args[1]
	serverName := "8.8.8.8:53"
	if len(os.Args) == 3 {
		serverName = os.Args[2]
	}
	ips, err := client.ResolveDNS(hostname, serverName)
	if err != nil {
		fmt.Printf("Failed to resolve DNS: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("IP addresses for %s:\n", hostname)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
