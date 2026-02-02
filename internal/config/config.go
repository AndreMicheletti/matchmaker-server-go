package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	POSTGRES_HOST string
	REDIS_HOST string
	REDIS_USER string
	REDIS_PASS string
}

func Load() Config {
	cfg := Config{}
	dotenv_file := os.Getenv("DOTENV")
	if dotenv_file == "" {
		dotenv_file = ".env.dev"
	}
	err := godotenv.Load(dotenv_file)
	if err != nil {
		log.Printf("%s\n", err)
		log.Fatal("Error loading .env file")
	}

	cfg.PORT = "8080"
	cfg.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	cfg.REDIS_HOST = os.Getenv("REDIS_HOST")
	cfg.REDIS_USER = os.Getenv("REDIS_USER")
	cfg.REDIS_PASS = os.Getenv("REDIS_PASS")
	return cfg
}
