package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort            string
	AppEnv             string
	DBPath             string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	FrontendURL        string
	// CORS
	AllowedOrigins []string
	// Email configuration
	SMTPHost  string
	SMTPPort  int
	SMTPUser  string
	SMTPPass  string
	FromEmail string
	FromName  string
	// Session
	SessionTTL time.Duration

	// Argon2id
	Argon2Memory     uint32 // KiB
	Argon2Iterations uint32
	Argon2Threads    uint8
}

var AppConfig *Config

func Load() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		slog.Warn("no .env file found, using environment variables")
	}

	AppConfig = &Config{
		AppPort:            getEnv("APP_PORT", "8080"),
		AppEnv:             getEnv("APP_ENV", "development"),
		DBPath:             getEnv("DB_PATH", "./data/app.db"),
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectURL:  getEnv("GOOGLE_REDIRECT_URL", ""),
		FrontendURL:        getEnv("FRONTEND_URL", "http://localhost:5173"),
		AllowedOrigins:     parseAllowedOrigins(),
		// Email configuration
		SMTPHost:  getEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:  getEnvAsInt("SMTP_PORT", 587),
		SMTPUser:  getEnv("SMTP_USER", ""),
		SMTPPass:  getEnv("SMTP_PASS", ""),
		FromEmail: getEnv("FROM_EMAIL", "noreply@example.com"),
		FromName:  getEnv("FROM_NAME", "Laju"),
		// Session
		SessionTTL: getSessionTTL(),

		// Argon2id
		Argon2Memory:     getArgon2Memory(),
		Argon2Iterations: getArgon2Iterations(),
		Argon2Threads:    getArgon2Threads(),
	}

	return AppConfig
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	var result int
	if _, err := fmt.Sscanf(value, "%d", &result); err != nil {
		return defaultValue
	}
	return result
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("file:%s?_foreign_keys=on", c.DBPath)
}

func (c *Config) IsDevelopment() bool {
	return c.AppEnv == "development"
}

// parseAllowedOrigins parses ALLOWED_ORIGINS env var (comma-separated).
// Defaults to http://localhost:5173 in dev, empty (strict) in prod.
func parseAllowedOrigins() []string {
	val := os.Getenv("ALLOWED_ORIGINS")
	if val == "" {
		return []string{"http://localhost:5173"}
	}
	var origins []string
	for _, o := range strings.Split(val, ",") {
		trimmed := strings.TrimSpace(o)
		if trimmed != "" {
			origins = append(origins, trimmed)
		}
	}
	return origins
}

func getSessionTTL() time.Duration {
	val := getEnv("SESSION_TTL", "168h")
	d, err := time.ParseDuration(val)
	if err != nil {
		return 24 * time.Hour
	}
	return d
}

// getArgon2Memory returns the argon2id memory cost (KiB) from env.
// Default: 65536 (64MB). Min: 1024 (1MB). Max: 1048576 (1GB).
func getArgon2Memory() uint32 {
	val := getEnv("ARGON2_MEMORY", "65536")
	v, err := strconv.ParseUint(val, 10, 32)
	if err != nil || v < 1024 || v > 1048576 {
		return 65536
	}
	return uint32(v)
}

// getArgon2Iterations returns the argon2id time cost from env.
// Default: 3. Min: 1. Max: 100.
func getArgon2Iterations() uint32 {
	val := getEnv("ARGON2_TIME", "3")
	v, err := strconv.ParseUint(val, 10, 32)
	if err != nil || v < 1 || v > 100 {
		return 3
	}
	return uint32(v)
}

// getArgon2Threads returns the argon2id parallelism from env.
// Default: 4. Min: 1. Max: 256.
func getArgon2Threads() uint8 {
	val := getEnv("ARGON2_THREADS", "4")
	v, err := strconv.ParseUint(val, 10, 8)
	if err != nil || v < 1 || v > 256 {
		return 4
	}
	return uint8(v)
}
