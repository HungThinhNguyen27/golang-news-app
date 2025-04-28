package elasticsearch

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

// ArticleStorage handles Elasticsearch operations for articles.
type ArticleStorage struct {
	esClient *elasticsearch.Client
}

// InitElasticsearch initializes the Elasticsearch client and returns a storage handler.
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
