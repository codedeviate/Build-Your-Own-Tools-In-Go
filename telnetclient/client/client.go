// client.go
package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// ConnectToServer connects to the Telnet server and handles communication.
func ConnectToServer(address string) error {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        return err
    }
    defer conn.Close()

    go func() {
        scanner := bufio.NewScanner(conn)
        for scanner.Scan() {
            fmt.Println(scanner.Text())
        }
    }()

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        _, err := fmt.Fprintf(conn, "%s\n", scanner.Text())
        if err != nil {
            return err
        }
    }

    return scanner.Err()
}
