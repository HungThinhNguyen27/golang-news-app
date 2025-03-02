package crawler

import (
	"crawl-service/config"
	"crawl-service/models"
	postgres "crawl-service/storage/postgres"
	"crypto/md5"
	"encoding/hex"
	"log"
)

func generateMD5(content string) string {
	hash := md5.Sum([]byte(content))
	return hex.EncodeToString(hash[:])
}

func CrawlArticles() []models.Article {

	var articles []models.Article

	// connect to DB on docker
	db := postgres.ConnectToDB()
	postgres.CreateArticlesTable(db)

	// crawl category URL in Vnexpress
	categoryURLs := FetchCategories(config.BASE_URL, config.ALLOWED_DOMAINS)
	for _, categoryURL := range categoryURLs {
		articleURLs := FetchArticlesURL(categoryURL, config.ALLOWED_DOMAINS) // get article urls in category page
		for _, articleURL := range articleURLs {
			articleDetail := FetchArticleDetail(articleURL, config.ALLOWED_DOMAINS) // get article detail like (titele, category, sub-category, published date, conten ,description ....)
			if articleDetail.Title == "" || articleDetail.Content == "" || articleDetail.Category == "" || articleDetail.PublishedDate == "" {
				log.Println("skip article:", articleURL)
				continue
			}
			newHash := generateMD5(articleDetail.Content)
			checkExishHash := postgres.CheckHashExists(db, newHash) // check in db old hash compare new hash
			if checkExishHash {
				log.Println("Duplicate article")
				continue

			}
			article := models.Article{
				Title:         articleDetail.Title,
				Description:   articleDetail.Description,
				Category:      articleDetail.Category,
				SubCategory:   articleDetail.SubCategory,
				URL:           articleDetail.URL,
				PublishedDate: articleDetail.PublishedDate,
				ImageURL:      articleDetail.ImageURL,
				Content:       articleDetail.Content,
				Hash:          newHash,
			}
			articles = append(articles, article)
		}
	}
	return articles // Return all collected articles
}
