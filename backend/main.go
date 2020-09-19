package main

import (
	"bytes"
	"log"
	"os"

	_ "github.com/lib/pq"
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

	file, err := os.Open(os.Getenv("SQL_INIT_PATH"))
	if err != nil {
		log.Fatalf("Failed to open sql init path: %s, error: %s\n", os.Getenv("SQL_INIT_PATH"), err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	query := buf.String()
	_, err = a.Client.DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to execute sql init file: %s\n", err)
		os.Exit(1)
	} else {
		log.Println("Successfully executed sql init file")
	}

	a.Run(":8080")
}
