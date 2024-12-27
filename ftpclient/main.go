// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/ftpclient/client"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: ftpclient <server> <user> <password> <filePath>")
		os.Exit(1)
	}

	server := os.Args[1]
	user := os.Args[2]
	password := os.Args[3]
	filePath := os.Args[4]

	content, err := client.FetchFile(server, user, password, filePath)
	if err != nil {
		fmt.Printf("Failed to fetch file %s: %v\n", filePath, err)
		os.Exit(1)
	}

	fmt.Printf("Content of %s:\n%s\n", filePath, content)
}
