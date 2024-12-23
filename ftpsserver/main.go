// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/ftpsserver/servermod"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: ftpsserver <port> <path> <certfile> <keyfile>")
		os.Exit(1)
	}

	port := os.Args[1]
	path := os.Args[2]
	certfile := os.Args[3]
	keyfile := os.Args[4]

	fmt.Printf("Starting FTP server on port %s for %s...\n", port, path)
	if err := servermod.StartServer(port, path, certfile, keyfile); err != nil {
		fmt.Printf("Server failed: %v\n", err)
		os.Exit(1)
	}
}
