package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	if err := godotenv.Load("config.env"); err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "tlaloc_security"),
		DBPort:     getEnv("DB_PORT", "5432"),
		JWTSecret:  getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-this-in-production"),
		JWTRefresh: getEnv("JWT_REFRESH_SECRET", "your-super-secret-refresh-key-change-this-in-production"),
		ServerPort: getEnv("SERVER_PORT", "8081"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
