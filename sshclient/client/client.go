// client.go
package client

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

// SSHClient represents an SSH client.
type SSHClient struct {
	client *ssh.Client
}

// NewSSHClient creates a new SSH client.
func NewSSHClient(user, password, host string, port int) (*SSHClient, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	address := fmt.Sprintf("%s:%d", host, port)
	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, err
	}

	return &SSHClient{client: client}, nil
}

// RunCommand runs a command on the remote SSH server.
func (s *SSHClient) RunCommand(cmd string) (string, error) {
	session, err := s.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// Close closes the SSH client connection.
func (s *SSHClient) Close() error {
	return s.client.Close()
}
