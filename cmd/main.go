package main

import (
	"ddos/config"
	router "ddos/internal/api"
	"ddos/internal/api/handler"
	"ddos/internal/generator"
	"log"
)

func main() {
	cfg := config.Load()

	handler := handler.NewHandler(&generator.Generator{})
	router := router.NewGin(handler)

	err := router.Run(cfg.HTTP_PORT)
	if err != nil {
		log.Fatalf("Error while serving at %v: %v", cfg.HTTP_PORT, err)
	}
}
