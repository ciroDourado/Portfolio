package main

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	RedisHost  string `env:"REDIS_HOST,required"`
	RedisPort  int    `env:"REDIS_PORT" envDefault:"6379"`
}

func (c *Config) GetRedisURL() string {
	return fmt.Sprintf(
		"%s:%d",
		c.RedisHost, c.RedisPort,
	)
}

func Load() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
