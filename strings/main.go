// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/strings/extractor"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	strings, err := extractor.ExtractStrings(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, s := range strings {
		fmt.Println(s)
	}
}
