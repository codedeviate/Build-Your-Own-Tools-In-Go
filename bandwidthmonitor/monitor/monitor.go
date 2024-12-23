// monitor.go
package monitor

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type BandwidthMonitor struct {
    device       string
    snapshotLen  int32
    promiscuous  bool
    timeout      time.Duration
    handle       *pcap.Handle
    packetCount  int
    byteCount    int
}

func NewBandwidthMonitor(device string, snapshotLen int32, promiscuous bool, timeout time.Duration) (*BandwidthMonitor, error) {
    handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
    if err != nil {
        return nil, err
    }

    return &BandwidthMonitor{
        device:      device,
        snapshotLen: snapshotLen,
        promiscuous: promiscuous,
        timeout:     timeout,
        handle:      handle,
    }, nil
}

func (bm *BandwidthMonitor) Start() {
    packetSource := gopacket.NewPacketSource(bm.handle, bm.handle.LinkType())
    for packet := range packetSource.Packets() {
        bm.packetCount++
        bm.byteCount += len(packet.Data())
    }
}

func (bm *BandwidthMonitor) Stop() {
    bm.handle.Close()
}

func (bm *BandwidthMonitor) Report() {
    fmt.Printf("Packets: %d, Bytes: %d\n", bm.packetCount, bm.byteCount)
}
