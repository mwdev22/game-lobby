package main

import (
	"context"
	"gateway/api"
	"gateway/clients/grpc"
	"log"
	"time"

	config "github.com/mwdev22/gocfg"
	"github.com/mwdev22/grpclib/grpcclient"
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
	gprcClient, err := grpcclient.NewClient(context.Background(), grpcclient.WithDefaultTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}

	playerClient := grpc.NewPlayerClient(gprcClient.Conn())

	app := api.New(cfg,
		api.WithPlayerClient(playerClient),
	)

	if err := app.Run(); err != nil {
		log.Fatalf("error while running app: %v", err)
	}
}
