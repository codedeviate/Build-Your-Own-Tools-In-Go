// services/fake_ssh.go
package services

import (
	"crypto/tls"
)

// FakeSSHService simulates an SSH service.
type FakeSSHService struct {
	*GenericService
}

// NewFakeSSHService returns a new FakeSSHService.
func NewFakeSSHService(customBanner string, useTLS bool, tlsConfig *tls.Config) *FakeSSHService {
	banner := customBanner
	if banner == "" {
		banner = "SSH-2.0-OpenSSH_7.4 (Fake)"
	}
	return &FakeSSHService{
		GenericService: NewGenericService("ssh", banner, useTLS, tlsConfig),
	}
}
