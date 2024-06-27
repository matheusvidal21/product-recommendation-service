package main

import (
	server "github.com/matheusvidal21/product-recommendation-service/application/handlers/server"
	"log"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("error starting server: %s", err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("error running server: %s", err)
	}
}
