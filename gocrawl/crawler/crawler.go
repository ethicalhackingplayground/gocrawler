package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Crawl the host
func Crawl(c *colly.Collector, url string) {

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		e.Request.Visit(e.Request.AbsoluteURL(link))
	})

	// The request of each link visisted
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})
	c.Visit(url)
}
