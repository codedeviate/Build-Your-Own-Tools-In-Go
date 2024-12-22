// whois.go
package whois

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// Whois performs a whois query for the given domain and returns the result.
func Whois(domain string) (string, error) {
	conn, err := net.DialTimeout("tcp", "whois.nic.google:43", time.Second*10)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\r\n", domain)
	scanner := bufio.NewScanner(conn)
	var result strings.Builder
	for scanner.Scan() {
		result.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return result.String(), nil
}
