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
	viper.SetDefault("DATABASE_URL", "postgres://postgres:D1sappearedV1ctory@localhost:5432/ai-response?sslmode=disable")

	viper.SetDefault("AI_SERVICE_URL", "http://localhost:8081/mock-ai-service")

	return &Config{
		DatabaseURL:  viper.GetString("DATABASE_URL"),
		AIServiceURL: viper.GetString("AI_SERVICE_URL"),
	}, nil
}
