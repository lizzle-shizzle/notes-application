package main

import (
	"log"
	"os"

	"github.com/lizzle-shizzle/notes-application/backend/api"
)

func main() {
	a := api.App{
		Client: &api.Client{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_DATABASE"),
		},
	}
	err := a.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize service: %s\n", err)
		os.Exit(1)
	}

	a.Run(":8080")
}
