// server.go
package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// handleConnection handles incoming Telnet connections.
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		conn.Write([]byte("> "))
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		input = strings.TrimSpace(input)
		if input == "exit" {
			return
		}
		response := fmt.Sprintf("You said: %s\n", input)
		conn.Write([]byte(response))
	}
}

// StartServer starts the Telnet server on the specified port.
func StartServer(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	defer listener.Close()
	fmt.Printf("Telnet server started on port %s\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handleConnection(conn)
	}
}
