package main

import (
	"L0/internal/config"
	"L0/internal/postgres"
	"context"
	"fmt"
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

	fmt.Println(connect)
}
