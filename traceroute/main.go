// main.go
package main

import (
	"fmt"
	"os"

	traceroute "github.com/username/traceroute/traceroutemod"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: traceroute <address>")
		os.Exit(1)
	}

	address := os.Args[1]
	maxHops := 30
	hops, err := traceroute.Traceroute(address, maxHops)
	if err != nil {
		fmt.Printf("Traceroute to %s failed: %v\n", address, err)
		os.Exit(1)
	}

	fmt.Printf("Traceroute to %s:\n", address)
	for _, hop := range hops {
		fmt.Println(hop)
	}
}
