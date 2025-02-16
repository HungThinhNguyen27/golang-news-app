package main

import (
	"crawl-service/config"
	"crawl-service/crawler"
	"crawl-service/models"
	"crawl-service/storage"
	postgres "crawl-service/storage/postgres"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
)

func generateMD5(content string) string {
	hash := md5.Sum([]byte(content))
	return hex.EncodeToString(hash[:])
}

func main() {

	// Load file .env
	env := config.LoadEnv()
	storage.InitExcelFile()

	// connect database
	db := postgres.ConnectToDB()
	postgres.CreateArticlesTable(db) // IF NOT EXIST

	// crawl data
	CategoriesURL := crawler.FetchCategories(env.BaseURL, env.AllowedDomains)
	for _, categoryURL := range CategoriesURL {
		articlesURL := crawler.FetchArticlesURL(categoryURL, env.AllowedDomains)
		for _, articleURL := range articlesURL { // syntax _ is index
			articleDetail := crawler.FetchArticleDetail(articleURL, env.AllowedDomains)
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
			postgres.SaveArticle(db, article)
			storage.SaveExcelFormat(article, articleURL)
		}
	}
	storage.SaveExcelFile(env.ExcelFile)
	fmt.Println("Crawl complete, data has been saved", env.ExcelFile)
}
