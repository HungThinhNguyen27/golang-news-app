package crawler

import (
	"crawl-service/config"
	"crawl-service/crawler"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test FetchCategories
func TestFetchCategories(t *testing.T) {
	result := crawler.FetchCategories(config.BASE_URL, config.ALLOWED_DOMAINS)
	assert.NotEmpty(t, result, "Categories should not be empty")
}

// Test FetchArticlesURL
func TestFetchArticlesURL(t *testing.T) {
	CategoryURL := "https://vnexpress.net/cong-nghe"
	result := crawler.FetchArticlesURL(CategoryURL, config.ALLOWED_DOMAINS)
	assert.NotEmpty(t, result, "Article URLs should not be empty")
}

// Test FetchArticleDetail
func TestFetchArticleDetail(t *testing.T) {

	articleURL := "https://vnexpress.net/ap-luc-de-nang-ukraine-sau-khau-chien-cua-ong-trump-zelensky-vnepre-4855410.html"
	result := crawler.FetchArticleDetail(articleURL, config.ALLOWED_DOMAINS)

	assert.NotEmpty(t, result.Title, "Title should not be empty")
	assert.NotEmpty(t, result.Description, "Description should not be empty")
	assert.NotEmpty(t, result.Category, "Category should not be empty")
	assert.NotEmpty(t, result.SubCategory, "SubCategory should not be empty")
	assert.NotEmpty(t, result.Content, "Content should not be empty")
	assert.NotEmpty(t, result.PublishedDate, "PublishedDate should not be empty")
	assert.NotEmpty(t, result.ImageURL, "ImageURL should not be empty")
}
