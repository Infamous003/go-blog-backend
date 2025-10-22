package config

import (
	"fmt"
	"os"
)

type Config struct {
	PublicHost string // where the API runs
	Port       string // API port
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func InitConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "9090"),
		DBUser:     getEnv("DB_USER", "infamous"),
		DBPassword: getEnv("DB_PASSWORD", "Getalife@03"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "go_blog"),
	}
}

func getEnv(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallback
}

func GetDBURL(cfg Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
}
