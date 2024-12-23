// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/httpsserver/server"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: httpsserver <port> <certFile> <keyFile>")
		os.Exit(1)
	}

	port := os.Args[1]
	certFile := os.Args[2]
	keyFile := os.Args[3]
	fmt.Printf("Starting server on port %s...\n", port)
	if err := server.StartServer(port, certFile, keyFile); err != nil {
		fmt.Printf("Server failed: %v\n", err)
		os.Exit(1)
	}
}
