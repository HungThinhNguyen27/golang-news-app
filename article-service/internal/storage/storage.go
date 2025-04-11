package storage

import "article-service/internal/models"

type StorageWithPostgres interface {
	GetArticleById(id int64) (models.Article, error)
	GetAllArticle(limit, offset int) ([]models.Article, error)
	GetTotalArticles() (int, error)
	UpdateArticle(id int64, updatedArticle models.Article) error
	DeleteArticle(id int64) error
	CreateArticle(article models.Article) (int64, error)
}

type StorageWithElasticSearch interface {
	GetAllArticles(limit int, offset int) ([]models.Article, error)
	GetTotalArticles() (int, error)
	GetArticlesByKeyWord(keyword string) ([]models.Article, error)
	GetArticlesByCategory(category string, limit int, offset int) ([]models.Article, error)
	// CreateArticle(article models.Article) (int64, error)
}
