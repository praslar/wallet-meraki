package deploy

import (
	"fmt"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

var Config AppConfig

type AppConfig struct {
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	Dbname     string `env:"DB_NAME"`

	Secret string `env:"SECRET"`
}

func SetEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err)
	}
	err = env.Parse(&Config)
	if err != nil {
		fmt.Println(err)
	}
}

func LoadEnv() AppConfig {
	return Config
}
