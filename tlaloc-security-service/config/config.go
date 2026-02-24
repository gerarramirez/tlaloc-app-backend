package config

import (
	"os"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	JWTSecret  string
	JWTRefresh string
	ServerPort string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "192.168.56.3"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "admin123456"),
		DBName:     getEnv("DB_NAME", "tlaloc_finance"),
		DBPort:     getEnv("DB_PORT", "5432"),
		JWTSecret:  getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-in-production"),
		JWTRefresh: getEnv("JWT_REFRESH_SECRET", "your-super-secret-refresh-key-change-in-production"),
		ServerPort: getEnv("SERVER_PORT", "8081"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
