// main.go
package main

import (
	"fmt"
	"os"

	netstat "github.com/username/netstat/netstatmod"
)

func main() {
	connections, err := netstat.Netstat()
	if err != nil {
		fmt.Printf("Netstat failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Active network connections:")
	for _, conn := range connections {
		fmt.Println(conn)
	}
}
