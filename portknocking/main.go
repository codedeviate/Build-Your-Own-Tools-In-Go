// main.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	portknocking "github.com/username/portknocking/knocking"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: portknocking <host> <delay_ms> <port1> <port2> ... <portN>")
		os.Exit(1)
	}

	host := os.Args[1]
	delayMs, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid delay: %v\n", err)
		os.Exit(1)
	}
	delay := time.Duration(delayMs) * time.Millisecond

	var ports []int
	for _, arg := range os.Args[3:] {
		port, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Invalid port: %v\n", err)
			os.Exit(1)
		}
		ports = append(ports, port)
	}

	if err := portknocking.Knock(host, ports, delay); err != nil {
		fmt.Printf("Port knocking failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Port knocking sequence completed.")
}
