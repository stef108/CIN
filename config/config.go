package config

import "os"

// AppConfig holds the application configuration
type AppConfig struct {
	Port    string
	AppName string
	Version string
}

// LoadConfig reads configuration from environment variables
func LoadConfig() AppConfig {
	return AppConfig{
		Port:    getEnv("PORT", "8080"),
		AppName: getEnv("APP_NAME", "DevOps-Tracker"),
		Version: getEnv("APP_VERSION", "1.0.0"),
	}
}

// Helper to read env var with default fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
