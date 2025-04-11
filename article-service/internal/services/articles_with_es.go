package services

import (
	"article-service/internal/models"
	"article-service/internal/storage"

	"fmt"
)

// ArticleService defines the service layer
type ArticleServiceWithES struct {
	storage storage.StorageWithElasticSearch
}

// NewArticleService creates a new instance of ArticleService
func NewArticleServiceWithES(s storage.StorageWithElasticSearch) *ArticleServiceWithES {
	return &ArticleServiceWithES{storage: s}
}

func (s *ArticleServiceWithES) GetByKeyWord(keyword string) ([]models.Article, error) {

	articles, err := s.storage.GetArticlesByKeyWord(keyword)
	if err != nil {
		return nil, fmt.Errorf("error fetching articles from Elasticsearch: %v", err)
	}
	return articles, nil
}

func (s *ArticleServiceWithES) GetAllArticles(limit int, page int, maxLimit int) ([]models.Article, int, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	} else if limit > maxLimit {
		limit = maxLimit
	}
	offset := (page - 1) * limit
	totalArticle, err := s.storage.GetTotalArticles()
	if err != nil {
		return nil, 0, 0, err
	}
	articles, err := s.storage.GetAllArticles(limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPages := (totalArticle + limit - 1) / limit

	return articles, totalArticle, totalPages, nil
}

func (s *ArticleServiceWithES) GetByCategory(category string, limit int, page int, maxLimit int) ([]models.Article, int, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	} else if limit > maxLimit {
		limit = maxLimit
	}
	offset := (page - 1) * limit
	totalArticle, err := s.storage.GetTotalArticles()
	if err != nil {
		return nil, 0, 0, err
	}
	articles, err := s.storage.GetArticlesByCategory(category, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPages := (totalArticle + limit - 1) / limit

	return articles, totalArticle, totalPages, nil
}
