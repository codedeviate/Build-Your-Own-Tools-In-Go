// main.go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/username/bandwidthmonitor/monitor"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: bandwidthmonitor <network_interface>")
		os.Exit(1)
	}

	device := os.Args[1]
	snapshotLen := int32(1024)
	promiscuous := false
	timeout := 30 * time.Second

	bm, err := monitor.NewBandwidthMonitor(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatalf("Error creating bandwidth monitor: %v", err)
	}

	go bm.Start()

	// Monitor for 10 seconds
	time.Sleep(10 * time.Second)

	bm.Stop()
	bm.Report()
}
