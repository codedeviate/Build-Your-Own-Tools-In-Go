// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/username/sshserver/server"
	"golang.org/x/crypto/ssh"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: sshserver <port> <private_key_file>")
		os.Exit(1)
	}

	port := os.Args[1]
	privateKeyFile := os.Args[2]

	privateBytes, err := os.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	config := &ssh.ServerConfig{
		NoClientAuth: true,
	}
	config.AddHostKey(private)

	fmt.Printf("Starting SSH server on port %s...\n", port)
	if err := server.StartServer(port, config); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
