package main

import (
	"article-service/internal/server"
	"log/slog"
	"os"
)

func main() {
	app, err := server.NewApp()
	if err != nil {
		slog.Error("Application failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
	app.StartServer()
}

// func main() {
//
// 	esStorage, err := elasticsearch.InitElasticsearch()
// 	if err != nil {
// 		fmt.Println("Error initializing Elasticsearch:", err)
// 		return
// 	}

//
// 	articleService := services.NewArticleServiceWithES(esStorage)
// 	fmt.Println("ArticleServiceWithES initialized:", articleService)

//
// 	articles, totalArticle, totalPages, err := articleService.GetAllArticles(10, 1, 30)
// 	fmt.Println("totalArticle:", totalArticle)
// 	fmt.Println("totalPages:", totalPages)

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

//
// 	for _, article := range articles {
// 		fmt.Println("Article:", article)
// 	}
// }
