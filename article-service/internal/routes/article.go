package routes

import (
	"article-service/internal/handlers/article"
	"article-service/internal/services"
	"net/http"

	"github.com/rs/cors"
)

// SetupRouter registers all API routes with CORS enabled
func SetupRouter(articleService *services.ArticleServiceWithES) http.Handler {
	articleHandler := article.NewArticleHandlerWithES(articleService)

	router := http.NewServeMux()
	router.HandleFunc("GET /articles", articleHandler.GetArticles)
	router.HandleFunc("GET /{category}", articleHandler.GetByCategory)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	return corsHandler
}
