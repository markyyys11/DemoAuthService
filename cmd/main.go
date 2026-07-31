package main

import (
	"DemoAuthService/internal/api"
	"DemoAuthService/internal/config"
	"DemoAuthService/internal/handlers/auth"
	"DemoAuthService/internal/logger"
	"DemoAuthService/internal/logger/loggers"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	logger.InitLogger(loggers.NewConsoleLogger())
	config.Load()

	_, err := pgx.Connect(context.Background(), config.Default().PostgresURL)
	if err != nil {
		log.Fatalf("Error connecting DB: %s", err.Error())
	}

	srv := api.NewServer()

	as := auth.NewService()
	ah := auth.NewHandler(as)

	srv.RegisterHandlers(ah)

	if err := srv.Run(); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}
