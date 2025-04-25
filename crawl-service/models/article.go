package models

type Article struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	SubCategory   string `json:"sub_category"`
	URL           string `json:"url"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Content       string `json:"content"`
	Hash          string `json:"hash"`
	Author        string `json:"author"`
}
