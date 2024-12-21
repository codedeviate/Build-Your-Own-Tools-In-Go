// sniffer.go
package sniffer

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"time"
)

// SniffPackets captures and prints packets from the specified network interface.
func SniffPackets(device string, snapshotLen int32, promiscuous bool, timeout time.Duration) error {
	handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		return err
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}

	return nil
}
