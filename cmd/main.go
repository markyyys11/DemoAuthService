package main

import (
	"DemoAuthService/internal/api"
	"DemoAuthService/internal/config"
	"DemoAuthService/internal/logger"
	"DemoAuthService/internal/logger/loggers"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	logger.InitLogger(loggers.NewConsoleLogger())
	config.Load()

	conn, err := pgxpool.New(context.Background(), config.Default().PostgresURL)
	if err != nil {
		log.Fatalf("Error connecting DB: %s", err.Error())
	}

	srv := api.NewServer(conn)
	if err := srv.Run(conn); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}
