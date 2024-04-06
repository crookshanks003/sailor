package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	Port     string
	User     string
	Password string
	DBname   string
}

type Config struct {
	Port     string
	DBConfig dbConfig
}

func ReadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %e", err)
	}

	return Config{
		Port: os.Getenv("PORT"),
		DBConfig: dbConfig{
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBname:   os.Getenv("DB_NAME"),
		},
	}
}
