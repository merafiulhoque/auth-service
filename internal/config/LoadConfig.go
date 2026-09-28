package config

import (
	"log/slog"
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
	err := godotenv.Load()

	if err != nil {
		slog.Error("ERROR: ", "err", err.Error())
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
