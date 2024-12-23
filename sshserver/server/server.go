// server.go
package server

import (
	"io"
	"log"
	"net"
	"strings"

	"golang.org/x/crypto/ssh"
)

// handleConnection handles incoming SSH connections.
func handleConnection(conn net.Conn, config *ssh.ServerConfig) {
	sshConn, chans, reqs, err := ssh.NewServerConn(conn, config)
	if err != nil {
		log.Printf("Failed to handshake: %v", err)
		return
	}
	defer sshConn.Close()

	go ssh.DiscardRequests(reqs)

	for newChannel := range chans {
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}

		channel, requests, err := newChannel.Accept()
		if err != nil {
			log.Printf("Could not accept channel: %v", err)
			return
		}

		go func(in <-chan *ssh.Request) {
			for req := range in {
				switch req.Type {
				case "shell":
					if len(req.Payload) == 0 {
						req.Reply(true, nil)
					}
				default:
					log.Printf("Command sent: %v\n", strings.Trim(string(req.Payload), "\x00\x0a\x0d"))
				}

			}
		}(requests)

		go func() {
			io.Copy(channel, channel)
			channel.Close()
		}()
	}
}

// StartServer starts the SSH server on the specified port.
func StartServer(port string, config *ssh.ServerConfig) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("SSH server started on port %s", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handleConnection(conn, config)
	}
}
