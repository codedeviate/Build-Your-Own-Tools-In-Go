// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/ftpserver/servermod"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ftpserver <port> <path>")
		os.Exit(1)
	}

	port := os.Args[1]
	path := os.Args[2]
	fmt.Printf("Starting FTP server on port %s for %s...\n", port, path)
	if err := servermod.StartServer(port, path); err != nil {
		fmt.Printf("Server failed: %v\n", err)
		os.Exit(1)
	}
}
