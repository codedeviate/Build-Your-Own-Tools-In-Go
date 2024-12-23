// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/smtpclient/client"
)

func main() {
	if len(os.Args) != 7 {
		fmt.Println("Usage: smtpclient <server> <port> <from> <to> <subject> <body>")
		os.Exit(1)
	}

	server := os.Args[1]
	port := os.Args[2]
	from := os.Args[3]
	to := os.Args[4]
	subject := os.Args[5]
	body := os.Args[6]

	err := client.SendEmail(server, port, from, to, subject, body)
	if err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Email sent successfully")
}
