package config

import "os"

type AppConfig struct {
	Port string
	Host string
}

func NewAppConfig() *AppConfig {
	cfg := &AppConfig{
		Port: LookUpEnv("APP_PORT", "8080"),
		Host: LookUpEnv("SENTRY_HOST", "example.com"),
	}
	return cfg
}

func LookUpEnv(key string, fallback string) string {
	if e, ok := os.LookupEnv(key); ok {
		return e
	} else {
		return fallback
	}
}
