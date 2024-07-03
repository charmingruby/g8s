package config

import (
	"log/slog"

	env "github.com/caarlos0/env/v6"
)

type environment struct {
	ServerPort string `env:"SERVER_PORT,required"`
}

func NewConfig() (*Config, error) {
	slog.Info("Loading environment...")
	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return nil, err
	}

	slog.Info("Environment loaded successfully!")

	cfg := Config{
		ServerConfig: &serverConfig{
			Port: environment.ServerPort,
		},
	}

	return &cfg, nil
}

type Config struct {
	ServerConfig *serverConfig
}

type serverConfig struct {
	Port string
}
