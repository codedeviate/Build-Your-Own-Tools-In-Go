// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/webcrawler/crawler"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: webcrawler <url>")
		os.Exit(1)
	}

	url := os.Args[1]
	fmt.Printf("Crawling URL: %s\n", url)
	crawler.Crawl(url)
}
