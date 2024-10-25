package main

import (
	"context"
	"os"
	"sila-app/internal/delivery"
	"sila-app/pkg/db/postgres"
	"sila-app/pkg/logger"
)

func main() {
	ctx := context.Background()
	mainLogger := logger.New("Sila")
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)

	postgresClient := postgres.New(ctx, os.Getenv("POSTGRES"))

	delivery.Start(postgresClient.DB, &mainLogger)
}
