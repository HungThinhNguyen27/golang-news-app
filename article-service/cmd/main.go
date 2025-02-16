package main

import (
	"article-service/handlers"
	"article-service/middleware"
	"article-service/repo"

	"log"
	"net/http"
)

func main() {
	// connect to postgres db
	db := repo.ConnectToDB()

	// Router and middleware
	mux := http.NewServeMux()
	mux.HandleFunc("/articles", handlers.GetArticles(db))

	log.Println("Server is running on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", middleware.EnableCORS(mux)))
}
