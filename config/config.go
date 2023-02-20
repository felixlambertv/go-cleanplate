package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var once sync.Once

type (
	Config struct {
		App
		HTTP
		Log
		PG
	}

	App struct {
		Name          string `env:"APP_NAME"`
		Version       string `env:"APP_VERSION"`
		Url           string `env:"APP_URL"`
		Secret        string `env:"APP_SECRET"`
		TokenLifespan int    `env:"TOKEN_HOUR_LIFESPAN"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env:"LOGGER_LOG_LEVEL"`
	}

	PG struct {
		PoolMax      int    `env:"PG_POOL_MAX"`
		Host         string `env:"PG_HOST"`
		User         string `env:"PG_USER"`
		Password     string `env:"PG_PASSWORD"`
		DatabaseName string `env:"PG_DBNAME"`
		Port         string `env:"PG_PORT"`
		SslMode      string `env:"PG_SSL_MODE"`
	}
)

func GetInstance() *Config {
	var configInstance *Config

	if configInstance == nil {
		once.Do(func() {
			confInstance, err := newConfig()

			if err != nil {
				log.Fatal(err)
			}

			configInstance = confInstance
		})
	}

	return configInstance
}

func newConfig() (*Config, error) {
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

func (pg PG) GetDbConnectionUrl() string {
	connectionUrl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		pg.Host, pg.User, pg.Password, pg.DatabaseName, pg.Port, pg.SslMode,
	)
	return connectionUrl
}
