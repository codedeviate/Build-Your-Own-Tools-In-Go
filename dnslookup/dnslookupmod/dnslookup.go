// dnslookup.go
package dnslookup

import (
	"net"
)

// DNSLookup performs a DNS lookup for the given domain and returns the IP addresses.
func DNSLookup(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, ip := range ips {
		result = append(result, ip.String())
	}
	return result, nil
}
