package postgres

import (
	"article-service/internal/models"
	"article-service/internal/utils"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

// OpenDB opens a connection to PostgreSQL
func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}
	return db, nil
}

// func ConnectToDB() (*Postgres, error) {
// 	env := configs.LoadEnv()

// 	dsn := fmt.Sprintf(
// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		env.POSTGRES_HOST, env.POSTGRES_PORT, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB,
// 	)
// 	var counts int
// 	maxRetries := 10

// 	for counts < maxRetries {
// 		connection, err := OpenDB(dsn)
// 		if err == nil {
// 			log.Println("Connected to PostgreSQL!")
// 			return &Postgres{Db: connection}, nil
// 		}
// 		log.Printf("Postgres not yet ready (attempt %d), retrying...", counts+1)
// 		counts++
// 		time.Sleep(2 * time.Second)
// 	}
// 	return nil, fmt.Errorf("failed to connect to database after %d attempts", maxRetries)
// }

// ConnectToDB initializes and returns a Postgres storage instance
func ConnectToDB() (*Postgres, error) {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		return nil, fmt.Errorf("DSN environment variable is not set")
	}

	var counts int
	for counts < 10 {
		connection, err := OpenDB(dsn)
		if err == nil {
			log.Println("Connected to PostgreSQL!")
			return &Postgres{Db: connection}, nil
		}

		log.Printf("Postgres not yet ready (attempt %d), retrying...", counts+1)
		counts++
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to database after multiple attempts")
}

// GetArticleById fetches an article by ID
func (p *Postgres) GetArticleById(id int64) (models.Article, error) {
	var article models.Article
	err := p.Db.QueryRow("SELECT id, title, description, category, sub_category, url, published_date, image_url, content FROM articlesTable WHERE id = $1", id).
		Scan(&article.Id, &article.Title, &article.Description, &article.Category, &article.SubCategory, &article.URL, &article.PublishedDate, &article.ImageURL, &article.Content)

	if err != nil {
		return models.Article{}, err
	}
	return article, nil
}

// GetAllArticle
func (p *Postgres) GetAllArticle(limit, offset int) ([]models.Article, error) {

	// Fetch articles with pagination
	rows, err := p.Db.Query(`
		SELECT id, title, description, category, sub_category, url, published_date, image_url, content 
		FROM articlesTable 
		ORDER BY published_date DESC 
		LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.Id, &article.Title, &article.Description, &article.Category, &article.SubCategory, &article.URL, &article.PublishedDate, &article.ImageURL, &article.Content)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

// Get count of articles
func (p *Postgres) GetTotalArticles() (int, error) {
	var total int
	err := p.Db.QueryRow("SELECT COUNT(*) FROM articlesTable").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

// DeleteArticle deletes an article by ID
func (p *Postgres) DeleteArticle(id int64) error {
	_, err := p.Db.Exec("DELETE FROM articlesTable WHERE id = $1", id)
	return err
}

// UpdateArticle updates an article by ID
func (p *Postgres) UpdateArticle(id int64, updatedArticle models.Article) error {
	query := `
		UPDATE articlesTable 
		SET title = $1, description = $2, category = $3, sub_category = $4, image_url = $5, content = $6 
		WHERE id = $7`
	_, err := p.Db.Exec(query, updatedArticle.Title, updatedArticle.Description, updatedArticle.Category, updatedArticle.SubCategory, updatedArticle.ImageURL, updatedArticle.Content, id)
	return err
}

// CreateArticle inserts a new article into the database
func (p *Postgres) CreateArticle(article models.Article) (int64, error) {
	article.PublishedDate = utils.GetCurrentTimestamp()
	query := `
		INSERT INTO articlesTable (title, description, category, sub_category, url, published_date, image_url, content) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var articleID int64
	err := p.Db.QueryRow(query, article.Title, article.Description, article.Category, article.SubCategory, article.URL, article.PublishedDate, article.ImageURL, article.Content).Scan(&articleID)
	if err != nil {
		return 0, err
	}
	return articleID, nil
}
