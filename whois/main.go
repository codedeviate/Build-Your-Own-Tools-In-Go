// main.go
package main

import (
	"fmt"
	"os"

	whois "github.com/username/whois/whoismod"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: whois <domain>")
		os.Exit(1)
	}

	domain := os.Args[1]
	result, err := whois.Whois(domain)
	if err != nil {
		fmt.Printf("Whois query for %s failed: %v\n", domain, err)
		os.Exit(1)
	}

	fmt.Printf("Whois result for %s:\n%s", domain, result)
}
