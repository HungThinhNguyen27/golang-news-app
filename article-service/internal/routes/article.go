package routes

import (
	"article-service/internal/handlers/article"
	"article-service/internal/services"
	"net/http"
)

// SetupRouter registers all API routes
func SetupRouter(articleService *services.ArticleServiceWithES) *http.ServeMux {
	articleHandler := article.NewArticleHandlerWithES(articleService)

	router := http.NewServeMux()
	router.HandleFunc("GET /api/article/{key_word}", articleHandler.GetByKeyWord)
	router.HandleFunc("GET /api/articles", articleHandler.GetAll)
	// router.HandleFunc("PUT /api/article/{id}", articleHandler.Update)
	// router.HandleFunc("DELETE /api/article/{id}", articleHandler.Delete)
	// router.HandleFunc("POST /api/article", articleHandler.Create)
	return router
}
