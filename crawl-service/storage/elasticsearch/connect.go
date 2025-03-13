package elasticsearch

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var ESClient *elasticsearch.Client

func InitElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	// Kiểm tra kết nối
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting Elasticsearch info: %s", err)
	}
	defer res.Body.Close()

	fmt.Println("Elasticsearch Connected!")
	ESClient = client
}

func GetESClient() *elasticsearch.Client {
	if ESClient == nil {
		InitElasticsearch()
	}
	return ESClient
}
