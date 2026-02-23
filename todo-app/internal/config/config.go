package config

import "os"

type Config struct {
	HTTPPort    string
	DatabaseUrl string
}

func Getenv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func Load() Config {
	return Config{
		HTTPPort:    Getenv("HTTP_PORT", "8080"),
		DatabaseUrl: Getenv("DATABASE_URL", ""),
	}
}
