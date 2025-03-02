package postgres

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func ConnectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	var counts int
	for {
		connection, err := OpenDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready, retrying...")
			counts++
		} else {
			log.Println("Connected to PostgreSQL!")
			return connection
		}

		if counts > 10 {
			log.Println("Could not connect to PostgreSQL:", err)
			return nil
		}
		log.Println("Backing off for 2 seconds...")
		time.Sleep(2 * time.Second)
	}
}
