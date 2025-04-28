package models

type Article struct {
	Id            string `json:"id"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Category      string `json:"category" validate:"required"`
	SubCategory   string `json:"sub_category"`
	URL           string `json:"url" validate:"required"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url" validate:"required"`
	Content       string `json:"content" validate:"required"`
	Hash          string `json:"hash" `
	Author        string `json:"author"`
}
