package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	APP_PORT          = "APP_PORT"
	ELASTICSEARCH_URL = "ELASTICSEARCH_URL"
)

type Config struct {
	AppPort          string
	ElasticsearchURL string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		AppPort:          getEnv(APP_PORT, "3000"),
		ElasticsearchURL: getEnv(ELASTICSEARCH_URL, "http://localhost:9200"),
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
