package main

import (
	"crawl-service/config"
	"crawl-service/crawler"
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

func main() {

	// connect database
	db := postgres.ConnectToDB()
	postgres.CreateArticlesTable(db) // IF NOT EXIST

	// crawl data
	CategoriesURL := crawler.FetchCategories(config.BASE_URL, config.ALLOWED_DOMAINS)
	for _, categoryURL := range CategoriesURL {
		articlesURL := crawler.FetchArticlesURL(categoryURL, config.ALLOWED_DOMAINS)
		for _, articleURL := range articlesURL { // syntax _ is index
			articleDetail := crawler.FetchArticleDetail(articleURL, config.ALLOWED_DOMAINS)
			// Check if there is missing data then ignore
			if articleDetail.Title == "" || articleDetail.Content == "" || articleDetail.Category == "" || articleDetail.PublishedDate == "" {
				log.Println("skip article:", articleURL)
				continue
			}
			newHash := generateMD5(articleDetail.Content)
			existHash := postgres.CheckHashExists(db, newHash)
			if existHash {
				log.Println("Duplicated article !")
				continue
			}
			article := models.Article{
				Title:         articleDetail.Title,
				Description:   articleDetail.Description,
				Category:      articleDetail.Category,
				SubCategory:   articleDetail.SubCategory,
				URL:           articleURL,
				PublishedDate: articleDetail.PublishedDate,
				ImageURL:      articleDetail.ImageURL,
				Content:       articleDetail.Content,
				Hash:          newHash,
			}
			err := postgres.SaveArticle(db, article)
			if err != nil {
				log.Println("Failed to save article:", err)
			} else {
				log.Println("Article saved successfully!")
			}
		}
	}
}
