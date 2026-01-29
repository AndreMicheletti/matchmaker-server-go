package config

import (
	"os"
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	POSTGRES_HOST string
	REDIS_HOST string
	DATABASE_NAME string
	DATABASE_USER string
	DATABASE_PASSWORD string
}

func Load() Config {
	cfg := Config{}
	dotenv_file := os.Getenv("DOTENV")
	if dotenv_file == "" {
		dotenv_file = ".env.dev"
	}
	err := godotenv.Load(dotenv_file)
	if err != nil {
		fmt.Printf("%s\n", err)
		log.Fatal("Error loading .env file")
	}

	cfg.PORT = "8080"
	cfg.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	cfg.REDIS_HOST = os.Getenv("REDIS_HOST")
	cfg.DATABASE_NAME = os.Getenv("DATABASE_NAME")
	cfg.DATABASE_USER = os.Getenv("DATABASE_USER")
	cfg.DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	return cfg
}
