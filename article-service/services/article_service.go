package services

import (
	"article-service/models"
	"article-service/repo"
	"database/sql"
)

func GetArticles(db *sql.DB) ([]models.Article, error) {
	return repo.GetAllArticles(db)
}
