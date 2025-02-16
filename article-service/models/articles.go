package models

type Article struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	SubCategory   string `json:"subCategory"`
	URL           string `json:"url"`
	PublishedDate string `json:"publishedDate"`
	ImageURL      string `json:"imageURL"`
	Content       string `json:"content"`
	Hash          string `json:"hash"`
}
