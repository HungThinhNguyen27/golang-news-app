package services

import (
	"article-service/internal/models"
	"article-service/internal/storage"
	"errors"
)

// ArticleService defines the service layer
type ArticleService struct {
	storage storage.Storage
}

// NewArticleService creates a new instance of ArticleService
func NewArticleService(s storage.Storage) *ArticleService {
	return &ArticleService{storage: s}
}

// GetPaginatedArticles fetches articles with pagination
func (s *ArticleService) GetPaginatedArticles(page, limit, maxLimit int) ([]models.Article, int, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	} else if limit > maxLimit {
		limit = maxLimit
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch articles from storage
	total, err := s.storage.GetTotalArticles()
	if err != nil {
		return nil, 0, 0, err
	}
	articles, err := s.storage.GetAllArticle(limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	// Calculate total pages
	totalPages := (total + limit - 1) / limit // Round up

	return articles, total, totalPages, nil
}

// GetArticleByID fetches an article by its ID
func (s *ArticleService) GetArticleByID(id int64) (models.Article, error) {
	// Validate ID
	if id <= 0 {
		return models.Article{}, errors.New("invalid article ID")
	}
	// Fetch article from storage
	article, err := s.storage.GetArticleById(id)
	if err != nil {
		return models.Article{}, errors.New("Article not found")
	}
	return article, nil
}

func (s *ArticleService) DeleteArticle(id int64) error {
	if id <= 0 {
		return errors.New("Invalid article ID")
	}
	_, err := s.GetArticleByID(id)
	if err != nil {
		return errors.New("Article not found")
	}
	return s.storage.DeleteArticle(id)
}

func (s *ArticleService) UpdateArticle(id int64, updatedArticle models.Article) error {
	if id <= 0 {
		return errors.New("invalid article ID")
	}
	_, err := s.GetArticleByID(id)
	if err != nil {
		return errors.New("Article not found")
	}
	return s.storage.UpdateArticle(id, updatedArticle)
}

func (s *ArticleService) CreateArticle(article models.Article) (int64, error) {
	// Store the article and get the inserted ID
	articleID, err := s.storage.CreateArticle(article)
	if err != nil {
		return 0, err
	}

	return articleID, nil
}
