package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type WorkerConfig struct {
	RedisHost  string `env:"REDIS_HOST,required"`
	RedisPort  int    `env:"REDIS_PORT" envDefault:"6379"`
}

func (wc *WorkerConfig) GetRedisURL() string {
	return fmt.Sprintf(
		"%s:%d",
		wc.RedisHost, wc.RedisPort,
	)
}

func LoadWorkerConfig() (*WorkerConfig, error) {
	var cfg WorkerConfig

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
