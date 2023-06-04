package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"log"
)

var Config AppConfig

type AppConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Dbname   string `env:"DB_NAME"`
}

func SetEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = env.Parse(&Config)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadEnv() AppConfig {
	return Config
}
