package config

import (
	"os"
)

type AppConfig struct {
	Name        string
	Version     string
	Environment string
	LogLevel    string
}

func NewAppConfig() AppConfig {
	return AppConfig{
		Name:        getEnvOrDefault("APP_NAME", "Amethis"),
		Version:     getEnvOrDefault("APP_VERSION", "1.0.0"),
		Environment: getEnvOrDefault("APP_ENV", "development"),
		LogLevel:    getEnvOrDefault("LOG_LEVEL", "info"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
