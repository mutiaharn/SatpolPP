package config

import (
	"os"
)

type Config struct {
	AppPort     string
	AppEnv      string
	JWTSecret   string
	DBDriver    string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DBName      string
	StoragePath string
}

func LoadConfig() *Config {
	return &Config{
		AppPort:     getEnv("APP_PORT", "8080"),
		AppEnv:      getEnv("APP_ENV", "development"),
		JWTSecret:   getEnv("JWT_SECRET", "super-secret-satpolpp-sulsel-diskominfo-2026"),
		DBDriver:    getEnv("DB_DRIVER", "sqlite"), // sqlite for easy dev, postgres for Diskominfo prod
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPass:      getEnv("DB_PASS", "postgres"),
		DBName:      getEnv("DB_NAME", "satpolpp_sulsel_db"),
		StoragePath: getEnv("STORAGE_PATH", "./uploads"),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
