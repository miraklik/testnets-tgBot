package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Tokens struct {
		Token string
	}

	MongoDB struct {
		URI string
	}
}

func LoadConfig() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var cfg config

	cfg.Tokens.Token = os.Getenv("TG_TOKEN")
	cfg.MongoDB.URI = os.Getenv("MONGO_URI")

	return &cfg, nil
}
