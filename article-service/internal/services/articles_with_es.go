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

// func (s *ArticleServiceWithES) GetByKeyWord(keyword string, limit int, offset int) ([]models.Article, error) {

// 	articles, err := s.storage.GetArticlesByKeyWord(keyword, limit, offset)
// 	if err != nil {
// 		return nil, fmt.Errorf("error fetching articles from Elasticsearch: %v", err)
// 	}
// 	return articles, nil
// }

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
	totalArticle, err := s.storage.CountArticles("", "")
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

func (s *ArticleServiceWithES) GetByCategory(
	category string,
	limit int,
	page int,
	maxLimit int,
) ([]models.Article, int, int, error) {

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	} else if limit > maxLimit {
		limit = maxLimit
	}
	offset := (page - 1) * limit
	totalArticle, err := s.storage.CountArticles("", category)
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

func (s *ArticleServiceWithES) GetByKeyWord(
	keyword string,
	limit int,
	page int,
	maxLimit int,
) ([]models.Article, int, int, error) {

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	} else if limit > maxLimit {
		limit = maxLimit
	}
	offset := (page - 1) * limit

	totalArticle, err := s.storage.CountArticles(keyword, "")
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to count articles: %v", err)
	}

	if totalArticle == 0 {
		return []models.Article{}, 0, 0, nil
	}

	articles, err := s.storage.GetArticlesByKeyWord(keyword, limit, offset)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to fetch articles: %v", err)
	}

	totalPages := (totalArticle + limit - 1) / limit

	return articles, totalArticle, totalPages, nil
}
