package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"rest_go_kv/pkg/logger"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

//DSN

type DbConfig struct {
	Dsn string
}

// Token

type AuthConfig struct {
	Secret string
}

// Чтение .env файла

func LoadConfig() *Config {
	logger.Debug("Attempting to load .env file")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file.")
		log.Println("Error loading .env file.")
	} else {
		logger.Info(".env file loaded successfully")
	}

	config := &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
	logger.Info("Config loaded: DSN and TOKEN retrieved from environment variables")

	return config
}
