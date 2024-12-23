// crawler.go
package crawler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Crawl fetches the HTML content of a URL and prints the links found.
func Crawl(url string) {
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

    doc.Find("a").Each(func(index int, item *goquery.Selection) {
        link, _ := item.Attr("href")
        fmt.Println(link)
    })
}
