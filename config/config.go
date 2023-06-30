package config

import (
	"os"
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
		LogFile:    getEnv("LOG_FILE", "../log/log.txt"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
