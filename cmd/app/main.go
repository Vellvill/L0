package main

import (
	"L0/internal/application"
	"L0/internal/config"
	"L0/internal/postgres"
	"context"
	"log"
)

func main() {

	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	connect, err := postgres.NewClient(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	app, err := application.NewApplication(config, connect)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Start()
	if err != nil {
		log.Fatal(err)
	}

}
