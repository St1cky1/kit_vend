package config

import (
	"os"
	"strconv"
)

type KitVendingAPI struct {
	CompanyId int
	Login     string
	Password  string
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Server struct {
	Port string
}

type Config struct {
	Server        Server
	Database      Database
	KitVendingAPI KitVendingAPI
	LogLevel      string
}

func Load() *Config {
	return &Config{
		Server: Server{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: Database{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "user"),
			Password: getEnv("DB_PASSWORD", "pass"),
			DBName:   getEnv("DB_NAME", "kit_vend"),
		},
		KitVendingAPI: KitVendingAPI{
			CompanyId: getEnvInt("KIT_COMPANY_ID", 380649),
			Login:     getEnv("KIT_LOGIN", "demotest"),
			Password:  getEnv("KIT_PASSWORD", "vendor734102"),
		},
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
