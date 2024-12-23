// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/telnetclient/client"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: telnetclient <address>")
		os.Exit(1)
	}

	address := os.Args[1]
	fmt.Printf("Connecting to Telnet server at %s...\n", address)
	if err := client.ConnectToServer(address); err != nil {
		fmt.Printf("Connection failed: %v\n", err)
		os.Exit(1)
	}
}
