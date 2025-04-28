package crawler

import (
	"crawl-service/config"
	"crawl-service/models"
	"crawl-service/storage/elasticsearch"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

func generateMD5(content string) string {
	hash := md5.Sum([]byte(content))
	return hex.EncodeToString(hash[:])
}

func generateID(title, date string) string {
	data := fmt.Sprintf("%s-%s", strings.ToLower(title), date)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func extractAuthor(content string) string {
	lines := strings.Split(strings.TrimSpace(content), "\n") // Split content into lines
	if len(lines) > 0 {
		return strings.TrimSpace(lines[len(lines)-1]) // Get the last line
	}
	return "Unknown" // If the author is not found
}

func CrawlArticles() []models.Article {

	var articles []models.Article

	// connect to DB on docker
	// db := postgres.ConnectToDB()
	// postgres.CreateArticlesTable(db)

	categoryURLs := FetchCategories(config.BASE_URL, config.ALLOWED_DOMAINS) // crawl category URL in Vnexpress
	for _, categoryURL := range categoryURLs {
		articleURLs := FetchArticlesURL(categoryURL, config.ALLOWED_DOMAINS) // get article urls in category page
		for _, articleURL := range articleURLs {
			articleDetail := FetchArticleDetail(articleURL, config.ALLOWED_DOMAINS) // get article detail like (titele, category, sub-category, published date, conten ,description ....)
			if articleDetail.Title == "" || articleDetail.Content == "" || articleDetail.Category == "" || articleDetail.PublishedDate == "" {
				log.Println("skip article:", articleURL)
				continue
			}
			newHash := generateMD5(articleDetail.Content)
			articleID := generateID(articleDetail.Title, articleDetail.PublishedDate)
			author := extractAuthor(articleDetail.Content) //

			// checkExishHash := postgres.CheckHashExists(db, newHash) // check in db old hash compare new hash
			checkExishHash, _ := elasticsearch.CheckHashExists(newHash) // check in db old hash compare new hash
			if checkExishHash {
				log.Println("Duplicate article :", articleDetail.Title)
				continue
			}

			article := models.Article{
				ID:            articleID,
				Title:         articleDetail.Title,
				Description:   articleDetail.Description,
				Category:      articleDetail.Category,
				SubCategory:   articleDetail.SubCategory,
				URL:           articleURL,
				PublishedDate: articleDetail.PublishedDate,
				ImageURL:      articleDetail.ImageURL,
				Content:       articleDetail.Content,
				Hash:          newHash,
				Author:        author,
			}
			articles = append(articles, article)
		}
	}
	return articles // Return all collected articles
}
