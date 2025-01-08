// firewall.go
package lib

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// Rule represents a firewall rule.
type Rule struct {
	SrcIP   net.IP
	DstIP   net.IP
	SrcPort layers.TCPPort
	DstPort layers.TCPPort
	Action  string // "allow" or "deny"
}

// Firewall represents a simple firewall.
type Firewall struct {
	rules []Rule
}

// NewFirewall creates a new Firewall.
func NewFirewall() *Firewall {
	return &Firewall{}
}

// AddRule adds a rule to the firewall.
func (fw *Firewall) AddRule(rule Rule) {
	fw.rules = append(fw.rules, rule)
}

// CheckPacket checks if a packet matches any rule.
func (fw *Firewall) CheckPacket(packet gopacket.Packet) bool {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	tcpLayer := packet.Layer(layers.LayerTypeTCP)

	if ipLayer != nil && tcpLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		tcp, _ := tcpLayer.(*layers.TCP)

		for _, rule := range fw.rules {
			if rule.SrcIP.Equal(ip.SrcIP) && rule.DstIP.Equal(ip.DstIP) &&
				rule.SrcPort == tcp.SrcPort && rule.DstPort == tcp.DstPort {
				return rule.Action == "allow"
			}
		}
	}
	return false
}

// StartFirewall starts the firewall on the specified interface.
func StartFirewall(iface string, fw *Firewall) {
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("Failed to open device: %v", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		if fw.CheckPacket(packet) {
			fmt.Printf("[%s] Packet allowed: From %s, To: %s\n", time.Now().UTC(), packet.NetworkLayer().NetworkFlow().Src().String(), packet.NetworkLayer().NetworkFlow().Dst().String())
		} else {
			fmt.Printf("[%s] Packet denied: From %s, To: %s\n", time.Now().UTC(), packet.NetworkLayer().NetworkFlow().Src().String(), packet.NetworkLayer().NetworkFlow().Dst().String())
		}
	}
}
