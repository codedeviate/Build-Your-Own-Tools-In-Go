// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/telnetserver/server"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: telnetserver <port>")
		os.Exit(1)
	}

	port := os.Args[1]
	fmt.Printf("Starting Telnet server on port %s...\n", port)
	if err := server.StartServer(port); err != nil {
		fmt.Printf("Server failed: %v\n", err)
		os.Exit(1)
	}
}
