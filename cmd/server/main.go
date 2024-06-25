package main

import (
	"github.com/gofiber/fiber/v2"
	recover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/matheusvidal21/product-recommendation-service/framework/config"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/logging"
	"log"
)

func main() {
	conf := config.LoadConfig()
	logger.Info("Starting server...")

	app := fiber.New()
	app.Use(recover.New())

	if err := app.Listen(conf.AppPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
