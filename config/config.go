package config

import (
	"github.com/caarlos0/env/v8"
	"log"
)

var Config AppConfig

type AppConfig struct {
	DBHost     string `env:"DB_HOST" `
	DBPort     string `env:"DB_PORT" `
	DBUsername string `env:"DB_USERNAME" `
	DBPassword string `env:"DB_PASSWORD" `
	Dbname     string `env:"DB_DBNAME" `
	Secret     string `env:"SECRET"`
}

func SetEnv() {

	err := env.Parse(&Config)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadEnv() AppConfig {
	return Config
}
