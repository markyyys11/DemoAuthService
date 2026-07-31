package main

import (
	"DemoAuthService/internal/api"
	"DemoAuthService/internal/api/module/auth"
	"DemoAuthService/internal/config"
	"DemoAuthService/internal/logger"
	"DemoAuthService/internal/logger/loggers"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	logger.InitLogger(loggers.NewConsoleLogger())
	config.Load()

	conn, err := pgx.Connect(context.Background(), config.Default().PostgresURL)
	if err != nil {
		log.Fatalf("Error connecting DB: %s", err.Error())
	}

	srv := api.NewServer()

	authService := auth.NewService(conn)
	authHandler := auth.NewHandler(authService)

	srv.RegisterHandlers(authHandler)

	if err := srv.Run(); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}
