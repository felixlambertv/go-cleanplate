package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App
		HTTP
		Log
		PG
	}

	App struct {
		Name    string `env:"APP_NAME"`
		Version string `env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env:"LOGGER_LOG_LEVEL"`
	}

	PG struct {
		PoolMax int    `env:"PG_POOL_MAX"`
		URL     string ``
	}
)

func NewConfig() (*Config, error) {
	config := &Config{}

	err := cleanenv.ReadConfig(".env", config)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}

	return config, err
}
