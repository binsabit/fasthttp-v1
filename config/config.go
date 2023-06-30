package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DB_DSN     string
	LogFile    string
}

func Configure() Config {
	return Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DB_DSN:     getEnv("DB_DSN", "postgres://yerdaulet:pa55word@localhost:5432/prosclad"),
		LogFile:    getEnv("LOG_FILE", "./log/log.txt"),
	}
}

func getEnv(key, fallback string) string {
	dir, err := os.Getwd()
	if err != nil {
		return fallback
	}
	err = godotenv.Load(filepath.Join(dir, "config", "config.env"))
	if err != nil {
		log.Println(err)
		return fallback

	}
	return os.Getenv(key)
}
