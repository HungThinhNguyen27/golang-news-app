package article

import (
	"article-service/internal/models"
	"article-service/internal/services"
	response "article-service/internal/utils"
	"log"
	"net/http"
	"strconv"
)

type ArticleHanddlerWithES struct {
	services *services.ArticleServiceWithES
}

func NewArticleHandlerWithES(s *services.ArticleServiceWithES) *ArticleHanddlerWithES {
	return &ArticleHanddlerWithES{services: s}
}

const maxLimit = 30

func parsePaginationParams(r *http.Request) (int, int) {
	page := 1
	limit := 10

	if p, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil && l > 0 && l <= maxLimit {
		limit = l
	}

	return page, limit
}

func (h *ArticleHanddlerWithES) GetByCategory(w http.ResponseWriter, r *http.Request) {

	page, limit := parsePaginationParams(r)
	categoryStr := r.PathValue("category")
	var (
		articles      []models.Article
		totalArticles int
		totalPages    int
		err           error
	)
	articles, totalArticles, totalPages, err = h.services.GetByCategory(categoryStr, limit, page, maxLimit)
	if err != nil {
		log.Printf("Error fetching articles: %v", err)
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}
	response.WriteJson(w, http.StatusOK, map[string]interface{}{
		"page":           page,
		"limit":          limit,
		"total_articles": totalArticles,
		"total_pages":    totalPages,
		"articles":       articles,
	})
}

func (h *ArticleHanddlerWithES) GetArticles(w http.ResponseWriter, r *http.Request) {
	page, limit := parsePaginationParams(r)
	keyWord := r.URL.Query().Get("keyword")
	categoryStr := r.URL.Query().Get("category")

	var (
		articles      []models.Article
		totalArticles int
		totalPages    int
		err           error
	)

	switch {
	case keyWord != "" && categoryStr != "":
		articles, totalArticles, totalPages, err = h.services.GetByKeywordAndCategory(keyWord, categoryStr, limit, page, maxLimit)
	case keyWord != "":
		articles, totalArticles, totalPages, err = h.services.GetByKeyWord(keyWord, limit, page, maxLimit)
	case categoryStr != "":
		articles, totalArticles, totalPages, err = h.services.GetByCategory(categoryStr, limit, page, maxLimit)
	default:
		articles, totalArticles, totalPages, err = h.services.GetAllArticles(limit, page, maxLimit)
	}

	if err != nil {
		log.Printf("Error fetching articles: %v", err)
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	response.WriteJson(w, http.StatusOK, map[string]interface{}{
		"page":           page,
		"limit":          limit,
		"total_articles": totalArticles,
		"total_pages":    totalPages,
		"articles":       articles,
	})
}
