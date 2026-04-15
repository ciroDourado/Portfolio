package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type WorkerConfig struct {
	DBHost     string   `env:"DB_HOST,required"`
	DBUser     string   `env:"DB_USER,required"`
	DBPassword string   `env:"DB_PASSWORD,required"`
	DBName     string   `env:"DB_NAME,required"`
	DBPort     int      `env:"DB_PORT" envDefault:"5432"`
	RedisHost  string `env:"REDIS_HOST,required"`
	RedisPort  int    `env:"REDIS_PORT" envDefault:"6379"`
}

func (wc *WorkerConfig) GetRedisURL() string {
	return fmt.Sprintf(
		"%s:%d",
		wc.RedisHost, wc.RedisPort,
	)
}

func (wc *WorkerConfig) GetDatabaseURL() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		wc.DBHost, wc.DBUser, wc.DBPassword, wc.DBName, wc.DBPort,
	)
}

func LoadWorkerConfig() (*WorkerConfig, error) {
	var cfg WorkerConfig

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
