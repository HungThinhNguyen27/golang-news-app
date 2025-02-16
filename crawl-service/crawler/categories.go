package crawler

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

// Function to crawl main categoryURL
func FetchCategories(baseURL string, allowedDomains string) []string {

	var categoriesURL []string
	c := NewCollector(allowedDomains)
	// get the list of categories URL
	c.OnHTML("nav.main-nav ul.parent li a", func(e *colly.HTMLElement) {
		categoryURL := e.Attr("href")

		if categoryURL != "" && !strings.Contains(categoryURL, "javascript") {
			fullURL := baseURL + categoryURL
			fmt.Println("Found category:", fullURL)
			categoriesURL = append(categoriesURL, fullURL) // Store category URL

		}

	})
	// Visit the home page to get the catalog
	err := c.Visit(baseURL)
	if err != nil {
		log.Println("fetchURL Error:", err)
	}
	return categoriesURL

}
