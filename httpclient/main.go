// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/httpclient/client"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: httpclient <url>")
		os.Exit(1)
	}

	url := os.Args[1]
	content, err := client.FetchURL(url)
	if err != nil {
		fmt.Printf("Failed to fetch URL %s: %v\n", url, err)
		os.Exit(1)
	}

	fmt.Printf("Content of %s:\n%s\n", url, content)
}
