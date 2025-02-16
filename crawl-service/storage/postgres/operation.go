package postgres

import (
	"crawl-service/models"
	"database/sql"
	"log"
)

func CreateArticlesTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS articlesTable (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		category TEXT NOT NULL,
		sub_category TEXT,
		url TEXT NOT NULL,
		published_date TEXT,
		image_url TEXT,
		content TEXT,
		hash TEXT
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Error when creating articles table", err)
	}
	log.Println("The articles table has been created")
}

func SaveArticle(db *sql.DB, article models.Article) {

	query := `
	INSERT INTO articlesTable (title, description, category, sub_category, url, published_date, image_url, content, hash)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`
	_, err := db.Exec(query, article.Title, article.Description, article.Category, article.SubCategory, article.URL, article.PublishedDate, article.ImageURL, article.Content, article.Hash)
	if err != nil {
		log.Println("Error when save data to articles table", err)
	}
	log.Println("---------------------")
	log.Println("saved successfully:", article.Title)

}

func CheckHashExists(db *sql.DB, hash string) bool {

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM articlesTable WHERE hash = $1);`
	err := db.QueryRow(query, hash).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
