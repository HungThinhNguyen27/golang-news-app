package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvConfig contains configuration variables
type EnvConfig struct {
	POSTGRES_HOST     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_PORT     string
}

// LoadEnv function is read environment variables from `.env` file
func LoadEnv() *EnvConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("not found .env file.")
	}

	return &EnvConfig{
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
	}
}
