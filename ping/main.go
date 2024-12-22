// main.go
package main

import (
    "fmt"
    "os"
    "github.com/username/ping/pingmod"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: ping <address>")
        os.Exit(1)
    }

    address := os.Args[1]
    duration, err := ping.Ping(address)
    if err != nil {
        fmt.Printf("Ping to %s failed: %v\n", address, err)
        os.Exit(1)
    }

    fmt.Printf("Ping to %s: %v\n", address, duration)
}
