package config

import (
	"fmt"
	"net/url"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppPort    int      `env:"APP_PORT" envDefault:"8080"`
	DBHost     string   `env:"DB_HOST,required"`
	DBUser     string   `env:"DB_USER,required"`
	DBPassword string   `env:"DB_PASSWORD,required"`
	DBName     string   `env:"DB_NAME,required"`
	DBPort     int      `env:"DB_PORT" envDefault:"5432"`
	StorageURL *url.URL `env:"STORAGE_URL,required"`
}

func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort,
	)
}

func Load() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("Falha ao carregar configurações: %w", err)
	}

	return &cfg, nil
}
