package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/username/hexdump/lib"
)

func main() {
	// Parse command-line arguments
	filePath := flag.String("file", "", "Path to the file to hexdump")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Usage: hexdump -file <path-to-file>")
		os.Exit(1)
	}

	// Open the file
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Perform the hexdump
	if err := lib.Dump(file, os.Stdout); err != nil {
		fmt.Printf("Error performing hexdump: %v\n", err)
		os.Exit(1)
	}
}
