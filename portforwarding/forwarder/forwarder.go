// forwarder.go
package forwarder

import (
	"io"
	"log"
	"net"
)

// forward handles the forwarding of data between the client and the target server.
func forward(src net.Conn, dst net.Conn) {
    defer src.Close()
    defer dst.Close()
    io.Copy(src, dst)
}

// handleConnection handles the incoming client connection and forwards data to the target server.
func handleConnection(clientConn net.Conn, target string) {
    defer clientConn.Close()

    serverConn, err := net.Dial("tcp", target)
    if err != nil {
        log.Printf("Failed to connect to target server: %v", err)
        return
    }
    defer serverConn.Close()

    go forward(clientConn, serverConn)
    forward(serverConn, clientConn)
}

// StartForwarder starts the port forwarder on the specified port and forwards traffic to the target server.
func StartForwarder(listenPort, target string) error {
    listener, err := net.Listen("tcp", ":"+listenPort)
    if err != nil {
        return err
    }
    defer listener.Close()

    log.Printf("Port forwarder started on port %s, forwarding to %s", listenPort, target)
    for {
        clientConn, err := listener.Accept()
        if err != nil {
            log.Printf("Failed to accept connection: %v", err)
            continue
        }
        go handleConnection(clientConn, target)
    }
}
