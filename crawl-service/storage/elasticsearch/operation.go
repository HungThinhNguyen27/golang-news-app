package elasticsearch

import (
	"bytes"
	"context"
	"crawl-service/models"
	"encoding/json"
	"fmt"
	"log"
)

// SaveDocument
func SaveDocument(article models.Article) error {
	esClient := GetESClient() // Get Elasticsearch client
	indexName := "articles"
	jsonData, err := json.Marshal(article)
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}

	// send request to Elasticsearch
	res, err := esClient.Index(
		indexName,
		bytes.NewReader(jsonData),
		esClient.Index.WithDocumentID(article.ID),
		esClient.Index.WithRefresh("true"),
	)
	if err != nil {
		return fmt.Errorf("error sending request to Elasticsearch: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing document in Elasticsearch: %s", res.String())
	}
	log.Printf("Document indexed successfully in %s with ID %s", indexName, article.Title)
	return nil
}

func CheckHashExists(hash string) (bool, error) {
	indexName := "articles"
	esClient := GetESClient() // Get Elasticsearch client

	query := []byte(fmt.Sprintf(`{"query": {"term": {"hash": "%s"}}}`, hash))
	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(indexName),
		esClient.Search.WithBody(bytes.NewReader(query)),
		esClient.Search.WithSize(1),
	)
	if err != nil || res.IsError() {
		return false, fmt.Errorf("Elasticsearch error: %v", err)
	}
	defer res.Body.Close()

	var result struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Hits.Total.Value > 0, nil
}
