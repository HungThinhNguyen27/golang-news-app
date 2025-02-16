package crawler

import (
	"log"

	"github.com/gocolly/colly/v2"
)

// Function to crawl url of articles in a category
func FetchArticlesURL(categoryURL string, allowedDomains string) []string {

	var urls []string
	c := NewCollector(allowedDomains)
	c.OnHTML("article.item-news", func(e *colly.HTMLElement) {
		articleURL := e.ChildAttr("h3.title-news a", "href")
		if articleURL != "" {
			urls = append(urls, articleURL)
		}
	})
	err := c.Visit(categoryURL)
	if err != nil {
		log.Println("CrawlArticles URL error:", categoryURL, "-", err)
	}
	return urls

}
