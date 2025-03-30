package article

import (
	"article-service/internal/services"
	"article-service/internal/utils/response"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type ArticleHanddlerWithES struct {
	services *services.ArticleServiceWithES
}

func NewArticleHandlerWithES(s *services.ArticleServiceWithES) *ArticleHanddlerWithES {
	return &ArticleHanddlerWithES{services: s}
}

func (h *ArticleHanddlerWithES) GetByKeyWord(w http.ResponseWriter, r *http.Request) {

	keyWord := r.PathValue("key_word")
	slog.Info("Fetching article", slog.String("key_word", keyWord))

	articles, err := h.services.GetByKeyWord(keyWord)
	if err != nil {
		slog.Error("Error fetching article", slog.String("key_word", keyWord))
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}
	// Return the article as JSON
	response.WriteJson(w, http.StatusOK, articles)
}

func (h *ArticleHanddlerWithES) GetAll(w http.ResponseWriter, r *http.Request) {
	maxLimit := 30
	page := 1
	limit := 10

	// Get 'page' from query parameter
	pageStr := r.URL.Query().Get("page")
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}
	// Get 'limit' from query parameter
	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil && l > 0 && l <= maxLimit {
			limit = l
		}
	}
	slog.Info("Get List article With:", slog.String("page", pageStr), slog.String("limit", limitStr))
	articles, totalArticles, totalPages, err := h.services.GetAllArticles(limit, page, maxLimit)
	if err != nil {
		log.Printf("Error fetching articles: %v", err)
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}
	// Return the article as JSON
	response.WriteJson(w, http.StatusOK, map[string]interface{}{
		"page":           page,
		"limit":          limit,
		"total_articles": totalArticles,
		"total_pages":    totalPages,
		"articles":       articles,
	})
}
