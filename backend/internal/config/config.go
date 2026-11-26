package config

import (
	"log/slog"
	"os"
)

type Config struct {
	AppPort          string
	AppEnv           string
	JWTSecret        string
	PIIEncryptionKey string
	DBDriver         string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPass           string
	DBName           string
	DBSSLMode        string
	StoragePath      string
	S3Endpoint       string
	S3Region         string
	S3AccessKey      string
	S3SecretKey      string
	S3BucketEvidence string
	S3BucketDocs     string
	S3BucketMedia    string
	S3UseSSL         bool
	SMTPHost         string
	SMTPPort         string
	SMTPUser         string
	SMTPPass         string
	SMTPFrom         string
	PublicSiteURL    string
}

func LoadConfig() *Config {
	cfg := &Config{
		AppPort:          getEnv("APP_PORT", "8080"),
		AppEnv:           getEnv("APP_ENV", "development"),
		JWTSecret:        getEnv("JWT_SECRET", "dev-only-jwt-secret-satpolpp-sulsel-2026-do-not-use-in-prod"),
		PIIEncryptionKey: getEnv("PII_ENCRYPTION_KEY", "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"),
		DBDriver:         getEnv("DB_DRIVER", "postgres"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "6432"),
		DBUser:           getEnv("DB_USER", "satpolpp_app"),
		DBPass:           getEnv("DB_PASS", "postgres"),
		DBName:           getEnv("DB_NAME", "satpolpp_sulsel_db"),
		DBSSLMode:        getEnv("DB_SSLMODE", "disable"),
		StoragePath:      getEnv("STORAGE_PATH", "./uploads"),
		S3Endpoint:       getEnv("S3_ENDPOINT", ""),
		S3Region:         getEnv("S3_REGION", "ap-southeast-1"),
		S3AccessKey:      getEnv("S3_ACCESS_KEY", ""),
		S3SecretKey:      getEnv("S3_SECRET_KEY", ""),
		S3BucketEvidence: getEnv("S3_BUCKET_EVIDENCE", "siaptertib-evidence"),
		S3BucketDocs:     getEnv("S3_BUCKET_DOCUMENTS", "portal-documents"),
		S3BucketMedia:    getEnv("S3_BUCKET_MEDIA", "portal-media"),
		S3UseSSL:         getEnv("S3_USE_SSL", "false") == "true",
		SMTPHost:         getEnv("SMTP_HOST", ""),
		SMTPPort:         getEnv("SMTP_PORT", "587"),
		SMTPUser:         getEnv("SMTP_USER", ""),
		SMTPPass:         getEnv("SMTP_PASS", ""),
		SMTPFrom:         getEnv("SMTP_FROM", "noreply@satpolpp.sulselprov.go.id"),
		PublicSiteURL:    getEnv("PUBLIC_SITE_URL", "http://localhost:3000"),
	}

	if cfg.AppEnv == "production" {
		if len(cfg.JWTSecret) < 32 {
			slog.Error("JWT_SECRET must be at least 32 characters in production")
			os.Exit(1)
		}
		if len(cfg.PIIEncryptionKey) != 64 {
			slog.Error("PII_ENCRYPTION_KEY must be a 64-char hex string (32 bytes) in production")
			os.Exit(1)
		}
	}

	return cfg
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
