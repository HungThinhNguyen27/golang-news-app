package elasticsearch

import (
	"article-service/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// GetArticlesByKeyword searches for articles by keyword in title or category fields.
func (s *ArticleStorage) GetArticlesByKeyword(keyword string, limit int, offset int) ([]models.Article, error) {
	query, err := buildKeywordQuery(keyword, limit, offset)
	if err != nil {
		return nil, err
	}
	return s.searchArticles(query)
}

// GetAllArticles retrieves all articles sorted by published date.
func (s *ArticleStorage) GetAllArticles(limit int, offset int) ([]models.Article, error) {
	query, err := buildMatchAllQuery(limit, offset)
	if err != nil {
		return nil, err
	}
	return s.searchArticles(query)
}

// CountArticles returns the number of articles matching a keyword and/or category.
func (s *ArticleStorage) CountArticles(keyword string, category string) (int, error) {
	query, err := buildCountQuery(keyword, category)
	if err != nil {
		return 0, err
	}

	res, err := s.esClient.Count(
		s.esClient.Count.WithContext(context.Background()),
		s.esClient.Count.WithIndex("articles"),
		s.esClient.Count.WithBody(bytes.NewReader(query)),
	)
	if err != nil {
		return 0, fmt.Errorf("error executing count query: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return 0, fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	var countResponse struct {
		Count int `json:"count"`
	}
	if err := json.NewDecoder(res.Body).Decode(&countResponse); err != nil {
		return 0, fmt.Errorf("error decoding count response: %w", err)
	}

	return countResponse.Count, nil
}

// GetArticlesByCategory retrieves articles filtered by category.
func (s *ArticleStorage) GetArticlesByCategory(category string, limit int, offset int) ([]models.Article, error) {
	query, err := buildCategoryQuery(category, limit, offset)
	if err != nil {
		return nil, err
	}
	return s.searchArticles(query)
}

// searchArticles is a helper to execute a search query and return parsed articles.
func (s *ArticleStorage) searchArticles(query []byte) ([]models.Article, error) {
	var articles []models.Article

	res, err := s.esClient.Search(
		s.esClient.Search.WithContext(context.Background()),
		s.esClient.Search.WithIndex("articles"),
		s.esClient.Search.WithBody(bytes.NewReader(query)),
		s.esClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("error executing search query: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	var esResponse struct {
		Hits struct {
			Hits []struct {
				Source models.Article `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&esResponse); err != nil {
		return nil, fmt.Errorf("error decoding search response: %w", err)
	}

	for _, hit := range esResponse.Hits.Hits {
		articles = append(articles, hit.Source)
	}
	return articles, nil
}

// GetArticlesByKeywordAndCategory retrieves articles filtered by both keyword and category.
func (s *ArticleStorage) GetArticlesByKeywordAndCategory(keyword, category string, limit int, offset int) ([]models.Article, error) {
	// Build the query for both keyword and category
	query, err := buildKeywordAndCategoryQuery(keyword, category, limit, offset)
	if err != nil {
		return nil, err
	}
	return s.searchArticles(query)
}
