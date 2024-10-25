package main

import (
	"context"
	"go-template/internal/delivery"
	"go-template/pkg/db/postgres"
	"go-template/pkg/logger"
	"os"
)

func main() {
	ctx := context.Background()
	mainLogger := logger.New("Sila")
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)

	postgresClient := postgres.New(os.Getenv("POSTGRES"))
	_ = postgresClient

	delivery.Start(postgresClient.DB, mainLogger)
}
