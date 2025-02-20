package routes

import (
	"article-service/internal/handlers/article"
	"article-service/internal/services"
	"net/http"
)

// SetupRouter registers all API routes
func SetupRouter(articleService *services.ArticleService) *http.ServeMux {
	articleHandler := article.NewArticleHandler(articleService)

	router := http.NewServeMux()
	router.HandleFunc("GET /api/article/{id}", articleHandler.GetByID)
	router.HandleFunc("GET /api/articles", articleHandler.GetList)
	router.HandleFunc("PUT /api/article/{id}", articleHandler.Update)
	router.HandleFunc("DELETE /api/article/{id}", articleHandler.Delete)
	router.HandleFunc("POST /api/article", articleHandler.Create)
	return router
}
