package main

import (
	"SecretSanta/pkg/config"
	"SecretSanta/pkg/services"
	"log"

	server "SecretSanta/pkg/rest"
)

func main() {
	appConfig := &config.Config{
		Host: "localhost",
		Port: "8080",
	}
	exchangeService := services.NewExchangeService()
	if err := startRESTEngine(appConfig, exchangeService); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func startRESTEngine(
	appConfig *config.Config,
	exchangeService *services.ExchangeService) error {

	restEngine := server.NewREST(appConfig, exchangeService)
	log.Printf("Starting server on %s:%s...", appConfig.Host, appConfig.Port)

	return restEngine.Run()
}
