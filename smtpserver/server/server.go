// server.go
package server

import (
	"errors"
	"io"
	"log"
	"time"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

// Backend implements SMTP server methods.
type Backend struct{}

// A Session is returned after successful login.
type Session struct{}

func (bkd *Backend) NewSession(conn *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

// AuthMechanisms returns a slice of available auth mechanisms; only PLAIN is
// supported in this example.
func (s *Session) AuthMechanisms() []string {
	return []string{sasl.Plain}
}

// Auth is the handler for supported authenticators.
func (s *Session) Auth(mech string) (sasl.Server, error) {
	return sasl.NewPlainServer(func(identity, username, password string) error {
		if username != "username" || password != "password" {
			return errors.New("Invalid username or password")
		}
		return nil
	}), nil
}

// Mail handles the MAIL command.
func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

// Rcpt handles the RCPT command.
func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Println("Rcpt to:", to)
	return nil
}

// Data handles the DATA command.
func (s *Session) Data(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	log.Println("Data received: ", string(data))
	return nil
}

// Reset handles the RSET command.
func (s *Session) Reset() {}

// Logout handles the LOGOUT command.
func (s *Session) Logout() error {
	return nil
}

// StartServer starts the SMTP server on the specified port.
func StartServer(port string) error {
	be := &Backend{}
	s := smtp.NewServer(be)

	s.Addr = ":" + port
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting SMTP server on port", port)
	return s.ListenAndServe()
}
