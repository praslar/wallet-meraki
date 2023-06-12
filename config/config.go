package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"log"
)

var Config AppConfig

type AppConfig struct {
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
<<<<<<< HEAD
	Dbname     string `env:"DB_DBNAME"`
	Secret     string `env:"SECRET"`
=======
	Dbname     string `env:"DB_NAME"`

	Secret string `env:"SECRET"`
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
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
