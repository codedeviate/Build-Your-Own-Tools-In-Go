// proxy.go
package proxy

import (
	"io"
	"log"
	"net"
)

// handleConnection handles the incoming client connection and forwards data to the target server.
func handleConnection(clientConn net.Conn, target string) {
    defer clientConn.Close()

    serverConn, err := net.Dial("tcp", target)
    if err != nil {
        log.Printf("Failed to connect to target server: %v", err)
        return
    }
    defer serverConn.Close()

    go io.Copy(serverConn, clientConn)
    io.Copy(clientConn, serverConn)
}

// StartProxy starts the proxy server on the specified port and forwards traffic to the target server.
func StartProxy(listenPort, target string) error {
    listener, err := net.Listen("tcp", ":"+listenPort)
    if err != nil {
        return err
    }
    defer listener.Close()

    log.Printf("Proxy server started on port %s, forwarding to %s", listenPort, target)
    for {
        clientConn, err := listener.Accept()
        if err != nil {
            log.Printf("Failed to accept connection: %v", err)
            continue
        }
        go handleConnection(clientConn, target)
    }
}
