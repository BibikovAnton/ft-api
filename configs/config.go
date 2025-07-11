package configs

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Sectet string
}

func LoadConfig() *Config {
	err := gotenv.Load()
	if err != nil {
		log.Println("Error loading .env file,use default config")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Sectet: os.Getenv("SECRET"),
		},
	}
}
