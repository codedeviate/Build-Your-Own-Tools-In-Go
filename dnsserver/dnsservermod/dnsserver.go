// dnsserver.go
package dnsserver

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

// handleDNSRequest handles incoming DNS requests.
func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	msg.Authoritative = true

	switch r.Opcode {
	case dns.OpcodeQuery:
		for _, q := range r.Question {
			switch q.Qtype {
			case dns.TypeA:
				ip := net.ParseIP("127.0.0.1")
				rr := &dns.A{
					Hdr: dns.RR_Header{
						Name:   q.Name,
						Rrtype: dns.TypeA,
						Class:  dns.ClassINET,
						Ttl:    0,
					},
					A: ip,
				}
				msg.Answer = append(msg.Answer, rr)
			}
		}
	}

	w.WriteMsg(&msg)
}

// StartDNSServer starts the DNS server on the specified port.
func StartDNSServer(port string) {
	server := &dns.Server{Addr: ":" + port, Net: "udp"}
	dns.HandleFunc(".", handleDNSRequest)
	log.Printf("Starting DNS server on port %s", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start DNS server: %v", err)
	}
}
