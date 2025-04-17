package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Config ...
type Config struct {
	Token string `env:"TOKEN,required"`
}

// NewConfig ...
func NewConfig(envFiles ...string) (*Config, error) {
	err := godotenv.Load(envFiles...)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}
	var config Config
	err = env.Parse(&config)
	if err != nil {
		return nil, fmt.Errorf("error parsing .env file: %w", err)
	}

	return &config, nil
}
