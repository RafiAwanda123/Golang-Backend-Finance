package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	DBUser        string
	DBPassword    string
	DBName        string
	DBHost        string
	DBPort        int
	ServerPort    int
	JWTSecret     string
	JWTExpiration time.Duration
}

// LoadConfig memuat konfigurasi dari environment variables
func LoadConfig() (*AppConfig, error) {
	cfg := &AppConfig{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnvAsInt("DB_PORT", 3307),
		ServerPort: getEnvAsInt("SERVER_PORT", 8080),
		JWTSecret:  getEnv("JWT_SECRET", "default-secret"),
	}

	// Validasi variabel wajib
	required := map[string]string{
		"DB_USER":     getEnv("DB_USER", ""),
		"DB_PASSWORD": getEnv("DB_PASS", ""),
		"DB_NAME":     getEnv("DB_NAME", ""),
	}

	for key, value := range required {
		if value == "" {
			return nil, fmt.Errorf("environment variable %s is required", key)
		}
	}

	cfg.DBUser = required["DB_USER"]
	cfg.DBPassword = required["DB_PASSWORD"]
	cfg.DBName = required["DB_NAME"]

	// Parse JWT expiration
	expirationStr := getEnv("JWT_EXP", "24h")
	expiration, err := time.ParseDuration(expirationStr)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXP format: %v", err)
	}
	cfg.JWTExpiration = expiration

	return cfg, nil
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	strValue := getEnv(key, "")
	if strValue == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(strValue)
	if err != nil {
		return defaultValue
	}
	return value
}
