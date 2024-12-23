// scraper.go
package scraper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Scrape fetches the HTML content of a URL and prints the text of the specified elements.
func Scrape(url string, selector string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Failed to fetch URL: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	doc.Find(selector).Each(func(index int, item *goquery.Selection) {
		text := item.Text()
		fmt.Println(text)
	})
}
