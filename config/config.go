package config

import (
	"os"
)

type Config struct {
	ServerPort string
	DB_DSN     string
}

func Configure() Config {
	return Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DB_DSN:     getEnv("DB_DSN", "postgres://yerdaulet:pa55word@localhost:5432/prosclad"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
