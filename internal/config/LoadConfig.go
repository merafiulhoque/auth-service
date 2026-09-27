package config

import (
	"auth-service/internal/shared/domainerrors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DB_URL        string
	JwtSecret     string
	RedisUrl      string
	ResendApiKey  string
	AllowedOrigin string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, domainerrors.ErrEnvLoad
	}

	return &Config{
		Port:          os.Getenv("PORT"),
		DB_URL:        os.Getenv("DB_URL"),
		AllowedOrigin: os.Getenv("ALLOWED_ORIGIN"),
		JwtSecret:     os.Getenv("JWT_SECRET"),
		RedisUrl:      os.Getenv("REDIS_URL"),
		ResendApiKey:  os.Getenv("RESEND_API_KEY"),
	}, nil
}
