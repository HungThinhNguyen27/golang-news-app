package elasticsearch

import (
	"article-service/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

type ArticleStorage struct {
	esClient *elasticsearch.Client
}

func InitElasticsearch() (*ArticleStorage, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating Elasticsearch client: %w", err)
	}

	res, err := client.Info()
	if err != nil {
		return nil, fmt.Errorf("error getting Elasticsearch info: %w", err)
	}
	defer res.Body.Close()

	fmt.Println("Elasticsearch Connected!")
	return &ArticleStorage{esClient: client}, nil
}

func (s *ArticleStorage) GetArticlesByKeyWord(keyword string, limit, offset int) ([]models.Article, error) {
	indexName := "articles"
	var articles []models.Article

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keyword,
				"fields": []string{"category", "title"},
			},
		},
		"from": offset,
		"size": limit,
		"sort": []map[string]interface{}{
			{"published_date.keyword": map[string]string{"order": "desc"}},
		},
	}

	queryBody, err := json.Marshal(searchQuery)
	if err != nil {
		return nil, fmt.Errorf("error marshalling query: %v", err)
	}

	res, err := s.esClient.Search(
		s.esClient.Search.WithContext(context.Background()),
		s.esClient.Search.WithIndex(indexName),
		s.esClient.Search.WithBody(bytes.NewReader(queryBody)),
		s.esClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("error executing search query: %v", err)
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
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	for _, hit := range esResponse.Hits.Hits {
		articles = append(articles, hit.Source)
	}

	return articles, nil
}

func (s *ArticleStorage) GetAllArticles(limit, offset int) ([]models.Article, error) {
	var articles []models.Article
	indexName := "articles"

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"from": offset,
		"size": limit,
		"sort": []map[string]interface{}{
			{"published_date.keyword": map[string]string{"order": "desc"}},
		},
	}

	queryBody, err := json.Marshal(searchQuery)
	if err != nil {
		return nil, fmt.Errorf("error marshalling query: %v", err)
	}

	res, err := s.esClient.Search(
		s.esClient.Search.WithContext(context.Background()),
		s.esClient.Search.WithIndex(indexName),
		s.esClient.Search.WithBody(bytes.NewReader(queryBody)),
		s.esClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("error executing search query: %v", err)
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
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	for _, hit := range esResponse.Hits.Hits {
		articles = append(articles, hit.Source)
	}

	return articles, nil
}

func (s *ArticleStorage) CountArticles(keyword string, category string) (int, error) {
	indexName := "articles"

	// Build dynamic bool query
	boolQuery := map[string]interface{}{}

	// match all if dont have filter
	if keyword == "" && category == "" {
		boolQuery = map[string]interface{}{
			"match_all": map[string]interface{}{},
		}
	} else {
		mustQueries := []map[string]interface{}{}

		if keyword != "" {
			mustQueries = append(mustQueries, map[string]interface{}{
				"multi_match": map[string]interface{}{
					"query":  keyword,
					"fields": []string{"title", "description", "content"},
				},
			})
		}

		if category != "" {
			mustQueries = append(mustQueries, map[string]interface{}{
				"match": map[string]interface{}{
					"category": category,
				},
			})
		}

		boolQuery = map[string]interface{}{
			"bool": map[string]interface{}{
				"must": mustQueries,
			},
		}
	}

	// Final count query
	countQuery := map[string]interface{}{
		"query": boolQuery,
	}

	queryBody, err := json.Marshal(countQuery)
	if err != nil {
		return 0, fmt.Errorf("error marshalling query: %v", err)
	}

	res, err := s.esClient.Count(
		s.esClient.Count.WithContext(context.Background()),
		s.esClient.Count.WithIndex(indexName),
		s.esClient.Count.WithBody(bytes.NewReader(queryBody)),
	)
	if err != nil {
		return 0, fmt.Errorf("error executing count query: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return 0, fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	var countResponse struct {
		Count int `json:"count"`
	}

	if err := json.NewDecoder(res.Body).Decode(&countResponse); err != nil {
		return 0, fmt.Errorf("error decoding response: %v", err)
	}
	return countResponse.Count, nil
}

func (s *ArticleStorage) GetArticlesByCategory(category string, limit, offset int) ([]models.Article, error) {
	var articles []models.Article
	indexname := "articles"

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{},
			},
		},
		"from": offset,
		"size": limit,
		"sort": []map[string]interface{}{
			{"published_date.keyword": map[string]string{"order": "desc"}},
		},
	}

	if category != "" {
		searchQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			searchQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"match": map[string]interface{}{"category": category},
			},
		)
	} else {
		searchQuery["query"] = map[string]interface{}{"match_all": map[string]interface{}{}}
	}

	queryBody, err := json.Marshal(searchQuery)
	if err != nil {
		return nil, fmt.Errorf("error marshalling query: %v", err)
	}

	res, err := s.esClient.Search(
		s.esClient.Search.WithContext(context.Background()),
		s.esClient.Search.WithIndex(indexname),
		s.esClient.Search.WithBody(bytes.NewReader(queryBody)),
		s.esClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("error executing search query: %v", err)
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
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	for _, hit := range esResponse.Hits.Hits {
		articles = append(articles, hit.Source)
	}

	return articles, nil
}
