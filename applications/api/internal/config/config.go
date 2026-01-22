package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppPort      string
	DatabaseURL  string
}

func Load() (*Config, error) {
	cfg := &Config{
		AppPort: getEnv("APP_PORT", "8080"),
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	cfg.DatabaseURL = dsn

	return cfg, nil
}

// Helper para ler env com fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}