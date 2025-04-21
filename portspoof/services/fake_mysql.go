// services/fake_mysql.go
package services

import (
	"crypto/tls"
)

// FakeMySQLService simulates a MySQL service.
type FakeMySQLService struct {
	*GenericService
}

// NewFakeMySQLService returns a new FakeMySQLService.
func NewFakeMySQLService(customBanner string, useTLS bool, tlsConfig *tls.Config) *FakeMySQLService {
	banner := customBanner
	if banner == "" {
		banner = "5.7.0-MySQL Community Server (Fake)"
	}
	return &FakeMySQLService{
		GenericService: NewGenericService("mysql", banner, useTLS, tlsConfig),
	}
}
