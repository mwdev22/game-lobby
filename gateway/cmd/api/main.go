package main

import (
	"gateway/api"
	"log"

	config "github.com/mwdev22/gocfg"
)

// @title           REST Boilerplate API
// @version         1.0
// @description     API documentation

func main() {
	cfg := config.New(
		config.WithDatabaseConfig(
			&config.DatabaseConfig{},
		),
	)
	app := api.New(cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("error while running app: %v", err)
	}
}
