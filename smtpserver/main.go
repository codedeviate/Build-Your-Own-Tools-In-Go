// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/smtpserver/server"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: smtpserver <port>")
		os.Exit(1)
	}

	port := os.Args[1]
	fmt.Printf("Starting SMTP server on port %s...\n", port)
	if err := server.StartServer(port); err != nil {
		fmt.Printf("Server failed: %v\n", err)
		os.Exit(1)
	}
}
