package article

import (
	"article-service/internal/models"
	service "article-service/internal/services"
	"article-service/internal/utils/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

// ArticleHandler handles HTTP requests for articles
type ArticleHandler struct {
	service *service.ArticleService
}

// NewArticleHandler creates a new ArticleHandler
func NewArticleHandler(s *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: s}
}

func (h *ArticleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// Extract the article ID from the URL path
	id := r.PathValue("id")
	slog.Info("Fetching article", slog.String("id", id))

	// Convert ID to int64
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		slog.Error("Invalid article ID", slog.String("id", id))
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	// Fetch article using service
	article, err := h.service.GetArticleByID(intID)
	if err != nil {
		slog.Error("Error fetching article", slog.String("id", id))
		response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
		return
	}

	// Return the article as JSON
	response.WriteJson(w, http.StatusOK, article)
}

// GetList handles the GET /api/articles request
func (h *ArticleHandler) GetList(w http.ResponseWriter, r *http.Request) {
	// Default values
	maxLimit := 100
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

	// Fetch paginated articles
	articles, total, totalPages, err := h.service.GetPaginatedArticles(page, limit, maxLimit)
	if err != nil {
		log.Printf("Error fetching articles: %v", err)
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	// Response with pagination metadata
	response.WriteJson(w, http.StatusOK, map[string]interface{}{
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
		"articles":    articles,
	})
}

func (h *ArticleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	slog.Info("Fetching article", slog.String("id", id))

	// Convert ID to int64
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		slog.Error("Invalid article ID", slog.String("id", id))
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}
	// Fetch article using service
	err = h.service.DeleteArticle(intID)
	if err != nil {
		slog.Error("Error fetching article", slog.String("id", id))
		response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
		return
	}

	// Return the article as JSON
	response.WriteJson(w, http.StatusOK, map[string]string{"message": "Article deleted successfully ", "article ID": id})
}

func (h *ArticleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	slog.Info("Update article ID ", slog.String("id", id))
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		slog.Error("Invalid article ID", slog.String("id", id))
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}
	var updatedArticle models.Article
	if err := json.NewDecoder(r.Body).Decode(&updatedArticle); err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}
	err = h.service.UpdateArticle(intID, updatedArticle)
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}

	response.WriteJson(w, http.StatusOK, map[string]string{"message": "Article updated successfully", "article ID": id})
}

func (h *ArticleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newArticle models.Article

	err := json.NewDecoder(r.Body).Decode(&newArticle)
	if errors.Is(err, io.EOF) {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
		return
	}
	articleID, err := h.service.CreateArticle(newArticle)
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}
	if err := validator.New().Struct(newArticle); err != nil {
		validateErrs := err.(validator.ValidationErrors)
		response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))

	}
	response.WriteJson(w, http.StatusCreated, map[string]interface{}{
		"message":    "Article created successfully",
		"article_id": articleID,
	})
}
