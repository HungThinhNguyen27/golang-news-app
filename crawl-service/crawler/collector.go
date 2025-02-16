package crawler

import (
	"github.com/gocolly/colly/v2"
)

// NewCollector creates a new Colly collector that is restricted
// to scrape only from the specified allowed domain.
func NewCollector(allowedDomains string) *colly.Collector {
	return colly.NewCollector(
		colly.AllowedDomains(allowedDomains),
	)
}
