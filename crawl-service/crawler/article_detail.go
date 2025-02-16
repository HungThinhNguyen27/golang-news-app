package crawler

import (
	"crawl-service/models"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func FetchArticleDetail(articleURL string, allowedDomains string) *models.Article {
	collector := NewCollector(allowedDomains)

	var title, description, category, subCategory, content, publishDate, imageURL string

	// extract article title
	collector.OnHTML("h1.title-detail", func(e *colly.HTMLElement) {
		title = strings.TrimSpace(e.Text)
	})
	// extract article description
	collector.OnHTML("p.description", func(e *colly.HTMLElement) {
		description = strings.TrimSpace(e.Text)
	})

	// extract article category example : Công nghệ
	collector.OnHTML("ul.breadcrumb li a", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			category = strings.TrimSpace(e.Text)
		}
	})

	// extract article sub-category example : AI -> công nghệ - AI
	collector.OnHTML("ul.breadcrumb li a", func(e *colly.HTMLElement) {
		if e.Index == 1 {
			subCategory = strings.TrimSpace(e.Text)
		}
	})

	// extract article content
	collector.OnHTML("article.fck_detail", func(e *colly.HTMLElement) {
		e.ForEach("p.Normal", func(_ int, el *colly.HTMLElement) {
			text := strings.TrimSpace(el.Text)
			if text != "" {
				content += text + "\n"
			}
		})
	})

	// extract article publishDate
	collector.OnHTML("span.date", func(e *colly.HTMLElement) {
		publishDate = strings.TrimSpace(e.Text)
	})

	// extract article imageURL
	collector.OnHTML("figure meta[itemprop='url']", func(e *colly.HTMLElement) {
		imageURL = e.Attr("content")
	})

	err := collector.Visit(articleURL)
	if err != nil {
		log.Println("Error visiting article:", err)
	}

	return &models.Article{
		Title:         title,
		Description:   description,
		Category:      category,
		SubCategory:   subCategory,
		Content:       content,
		PublishedDate: publishDate,
		ImageURL:      imageURL,
	}
}
