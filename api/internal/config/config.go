package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	Port   string
	AppEnv string
	DB     DBConfig
}

func Load() *Config {
	return &Config{
		Port:   getEnv("API_PORT", "8080"),
		AppEnv: getEnv("APP_ENV", "development"),
		DB: DBConfig{
			Host:     getEnv("PGHOST", "localhost"),
			Port:     getEnv("PGPORT", "5432"),
			User:     getEnv("PGUSER", "elogap"),
			Password: getEnv("PGPASSWORD", ""),
			Name:     getEnv("PGNAME", "elogap"),
		},
	}
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name,
	)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
