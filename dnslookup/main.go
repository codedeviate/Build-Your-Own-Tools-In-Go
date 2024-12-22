// main.go
package main

import (
	"fmt"
	"os"

	dnslookup "github.com/username/dnslookup/dnslookupmod"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: dnslookup <domain>")
		os.Exit(1)
	}

	domain := os.Args[1]
	ips, err := dnslookup.DNSLookup(domain)
	if err != nil {
		fmt.Printf("DNS lookup for %s failed: %v\n", domain, err)
		os.Exit(1)
	}

	fmt.Printf("DNS lookup result for %s:\n", domain)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
