package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func LoadConfig() (*Config, error) {
	// بارگذاری متغیرهای محیطی از فایل .env
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}

	return config, nil
}
