// client.go
package client

import (
	"net/smtp"
)

// SendEmail sends an email using the specified SMTP server.
func SendEmail(server, port, from, to, subject, body string) error {
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	addr := server + ":" + port

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()

	defer c.Quit()

	if err = c.Mail(from); err != nil {
		return err
	}

	if err = c.Rcpt(to); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
