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
	router.HandleFunc("GET /api/article/{key_word}", articleHandler.GetByKeyWord)
	router.HandleFunc("GET /api/articles", articleHandler.GetAll)

	// Thiết lập CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Cho phép frontend truy cập
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	return corsHandler
}
