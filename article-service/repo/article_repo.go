package repo

import (
	"article-service/models"
	"database/sql"
)

func GetAllArticles(db *sql.DB) ([]models.Article, error) {
	rows, err := db.Query("SELECT title, description, image_url, category, sub_category, url, published_date, content FROM articlesTable ORDER BY id desc LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.Title, &article.Description, &article.ImageURL, &article.Category, &article.SubCategory, &article.URL, &article.PublishedDate, &article.Content); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}
