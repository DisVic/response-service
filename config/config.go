package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL  string
	AIServiceURL string
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetDefault("DATABASE_URL", "postgres://postgres:D1sappearedV1ctory@localhost:5432/ai_response?sslmode=disable")
	viper.SetDefault("AI_SERVICE_URL", "http://aiservice/api")

	return &Config{
		DatabaseURL:  viper.GetString("DATABASE_URL"),
		AIServiceURL: viper.GetString("AI_SERVICE_URL"),
	}, nil
}
