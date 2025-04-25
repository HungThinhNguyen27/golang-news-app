package services

import (
	"article-service/internal/models"
	"article-service/internal/repo"
	"article-service/internal/utils"
	"fmt"
)

// ArticleServiceWithES defines the service layer using Elasticsearch as the storage
type ArticleServiceWithES struct {
	repository repo.ArticleRepository
}

// NewArticleServiceWithES creates a new instance of ArticleServiceWithES
func NewArticleServiceWithES(r repo.ArticleRepository) *ArticleServiceWithES {
	return &ArticleServiceWithES{repository: r}
}

// GetAllArticles returns all articles with pagination
func (r *ArticleServiceWithES) GetAllArticles(limit int, page int, maxLimit int) ([]models.Article, int, int, error) {
	limit, offset := utils.NormalizePagination(limit, page, maxLimit)

	totalArticle, err := r.repository.CountArticles("", "")
	if err != nil {
		return nil, 0, 0, err
	}

	articles, err := r.repository.GetAllArticles(limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := utils.CalculateTotalPages(totalArticle, limit)
	return articles, totalArticle, totalPages, nil
}

// GetByCategory filters articles by category with pagination
func (r *ArticleServiceWithES) GetByCategory(category string, limit int, page int, maxLimit int) ([]models.Article, int, int, error) {
	limit, offset := utils.NormalizePagination(limit, page, maxLimit)

	totalArticle, err := r.repository.CountArticles("", category)
	if err != nil {
		return nil, 0, 0, err
	}

	articles, err := r.repository.GetArticlesByCategory(category, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := utils.CalculateTotalPages(totalArticle, limit)
	return articles, totalArticle, totalPages, nil
}

// GetByKeyWord searches articles by keyword in title/category with pagination
func (r *ArticleServiceWithES) GetByKeyWord(keyword string, limit int, page int, maxLimit int) ([]models.Article, int, int, error) {
	limit, offset := utils.NormalizePagination(limit, page, maxLimit)

	totalArticle, err := r.repository.CountArticles(keyword, "")
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to count articles: %v", err)
	}

	if totalArticle == 0 {
		return []models.Article{}, 0, 0, nil
	}

	articles, err := r.repository.GetArticlesByKeyword(keyword, limit, offset)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to fetch articles: %v", err)
	}

	totalPages := utils.CalculateTotalPages(totalArticle, limit)
	return articles, totalArticle, totalPages, nil
}

// GetByKeywordAndCategory searches articles by keyword and filters by category with pagination
func (r *ArticleServiceWithES) GetByKeywordAndCategory(keyword string, category string, limit int, page int, maxLimit int) ([]models.Article, int, int, error) {
	limit, offset := utils.NormalizePagination(limit, page, maxLimit)

	totalArticle, err := r.repository.CountArticles(keyword, category)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to count articles by keyword and category: %v", err)
	}

	if totalArticle == 0 {
		return []models.Article{}, 0, 0, nil
	}

	articles, err := r.repository.GetArticlesByKeywordAndCategory(keyword, category, limit, offset)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to fetch articles by keyword and category: %v", err)
	}

	totalPages := utils.CalculateTotalPages(totalArticle, limit)
	return articles, totalArticle, totalPages, nil
}
