package config

import (
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

// Config holds application configuration values
type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = LoadConfig()

// LoadConfig initializes and returns the configuration
func LoadConfig() Config {
	godotenv.Load() // Load .env file

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", ""),
		Port:                   getEnv("PORT", "3000"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", ""),
		DBAddress:              getEnv("DB_HOST", "localhost"),
		DBName:                 getEnv("DB_NAME", "ecom"),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

// getEnv retrieves an environment variable or returns a fallback value
func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// getEnvAsInt retrieves an environment variable as an int64 or returns a fallback value
func getEnvAsInt(key string, fallback int64) int64 {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return fallback
	}
	return intValue
}
