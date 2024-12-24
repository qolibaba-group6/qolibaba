package config

import (
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
	DatabaseDSN   string
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No configuration file found: %v", err)
	}

	return Config{
		ServerAddress: viper.GetString("SERVER_ADDRESS"),
		DatabaseDSN:   viper.GetString("DATABASE_DSN"),
	}
}
