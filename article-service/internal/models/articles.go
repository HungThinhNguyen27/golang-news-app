package models

type Article struct {
	Id            int64  `json:"id"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Category      string `json:"category" validate:"required"`
	SubCategory   string `json:"subCategory"`
	URL           string `json:"url" validate:"required"`
	PublishedDate string `json:"publishedDate"`
	ImageURL      string `json:"imageURL" validate:"required"`
	Content       string `json:"content" validate:"required"`
	Hash          string `json:"hash" `
}
