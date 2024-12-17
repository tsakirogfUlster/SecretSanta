package main

import (
	"SecretSanta/pkg/config"
	"log"

	server "SecretSanta/pkg/rest"
)

func main() {
	appConfig := &config.Config{
		Host: "localhost",
		Port: "8080",
	}
	if err := startRESTEngine(appConfig); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func startRESTEngine(
	appConfig *config.Config) error {

	restEngine := server.NewREST(appConfig)

	return restEngine.Run()
}
