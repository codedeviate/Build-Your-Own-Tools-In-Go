// main.go
package main

import (
	"fmt"
	"os"

	dnsserver "github.com/username/dnsserver/dnsservermod"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: dnsserver <port>")
        os.Exit(1)
    }

    port := os.Args[1]

    fmt.Printf("Starting DNS server on port %s...\n", port)
    dnsserver.StartDNSServer(port)
}
