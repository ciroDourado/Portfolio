package config

import (
	"fmt"
	"net/url"

	"github.com/caarlos0/env/v11"
)

type ApiConfig struct {
	AppPort    int      `env:"APP_PORT" envDefault:"8080"`
	DBHost     string   `env:"DB_HOST,required"`
	DBUser     string   `env:"DB_USER,required"`
	DBPassword string   `env:"DB_PASSWORD,required"`
	DBName     string   `env:"DB_NAME,required"`
	DBPort     int      `env:"DB_PORT" envDefault:"5432"`
	RedisHost  string   `env:"REDIS_HOST,required"`
	RedisPort  int      `env:"REDIS_PORT" envDefault:"6379"`
	StorageURL *url.URL `env:"STORAGE_URL,required"`
}

func (ac *ApiConfig) GetDatabaseURL() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		ac.DBHost, ac.DBUser, ac.DBPassword, ac.DBName, ac.DBPort,
	)
}

func (ac *ApiConfig) GetRedisURL() string {
	return fmt.Sprintf(
		"%s:%d",
		ac.RedisHost, ac.RedisPort,
	)
}

func LoadApiConfig() (*ApiConfig, error) {
	var cfg ApiConfig

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("Falha ao carregar configurações: %w", err)
	}

	return &cfg, nil
}
