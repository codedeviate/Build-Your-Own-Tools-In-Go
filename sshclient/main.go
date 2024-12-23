// main.go
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/username/sshclient/client"
)

func main() {
	if len(os.Args) != 6 {
		fmt.Println("Usage: sshclient <user> <password> <host> <port> <command>")
		os.Exit(1)
	}

	user := os.Args[1]
	password := os.Args[2]
	host := os.Args[3]
	port, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Printf("Invalid port: %v\n", err)
		os.Exit(1)
	}
	command := os.Args[5]

	client, err := client.NewSSHClient(user, password, host, port)
	if err != nil {
		fmt.Printf("Failed to create SSH client: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	output, err := client.RunCommand(command)
	if err != nil {
		fmt.Printf("Failed to run command: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Command output:", output)
}
