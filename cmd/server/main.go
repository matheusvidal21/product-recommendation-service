package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	recover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/matheusvidal21/product-recommendation-service/framework/config"
	"github.com/matheusvidal21/product-recommendation-service/framework/database"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/logging"
	"log"
)

func main() {
	conf := config.LoadConfig()
	logger.Info("Starting server...")

	_, err := database.NewElasticsearchConnection(conf.ElasticsearchURL)
	if err != nil {
		log.Fatalf("Error connecting to Elasticsearch: %v", err)
	}
	logger.Info(fmt.Sprintf("Connected to Elasticsearch: %s", conf.ElasticsearchURL))

	app := fiber.New()
	app.Use(recover.New())

	if err := app.Listen(conf.AppPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
