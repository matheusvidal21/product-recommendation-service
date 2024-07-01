package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	APP_PORT     = "APP_PORT"
	POSTGRES_URL = "POSTGRES_URL"
	DB_NAME      = "DB_NAME"
	DB_USER      = "DB_USER"
	DB_PASSWORD  = "DB_PASSWORD"
	DB_PORT      = "DB_PORT"
)

type Config struct {
	AppPort     string
	PostgresURL string
	DBName      string
	DBUser      string
	DBPassword  string
	DBPort      string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		AppPort:     getEnv(APP_PORT, "3000"),
		PostgresURL: getEnv(POSTGRES_URL, "localhost"),
		DBName:      getEnv(DB_NAME, "product-recommendation"),
		DBUser:      getEnv(DB_USER, "postgres"),
		DBPassword:  getEnv(DB_PASSWORD, "postgres"),
		DBPort:      getEnv(DB_PORT, "5432"),
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
