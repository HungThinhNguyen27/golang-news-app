package repo

import "article-service/internal/models"

type ArticleRepository interface {
	GetArticlesByKeyword(keyword string, limit int, offset int) ([]models.Article, error)
	GetAllArticles(limit int, offset int) ([]models.Article, error)
	CountArticles(keyword string, category string) (int, error)
	GetArticlesByCategory(category string, limit int, offset int) ([]models.Article, error)
	GetArticlesByKeywordAndCategory(keyword, category string, limit int, offset int) ([]models.Article, error)
}
