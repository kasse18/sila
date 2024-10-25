package main

import (
	"context"
	"go-template/pkg/db/postgres"
	"go-template/pkg/logger"
	"os"
)

func main() {
	ctx := context.Background()
	mainLogger := logger.New("template")
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)

	postgresClient := postgres.New(os.Getenv("POSTGRES"))
	_ = postgresClient
}
