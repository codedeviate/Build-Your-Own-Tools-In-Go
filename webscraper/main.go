// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/webscraper/scraper"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: webscraper <url> <selector>")
		os.Exit(1)
	}

	url := os.Args[1]
	selector := os.Args[2]
	fmt.Printf("Scraping URL: %s with selector: %s\n", url, selector)
	scraper.Scrape(url, selector)
}
