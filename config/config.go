package config

import (
	"github.com/caarlos0/env/v8"
	"log"
)

var Config AppConfig

type AppConfig struct {
	DBHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUsername string `env:"DB_USERNAME" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"secret"`
	Dbname     string `env:"DB_DBNAME" envDefault:"postgres"`
	Secret     string `env:"SECRET" envDefault:"1234567890abcdef1234567890abcdef"`
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
