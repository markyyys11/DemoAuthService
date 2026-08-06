package api

import (
	"DemoAuthService/internal/api/middlewares"
	"DemoAuthService/internal/api/modules/auth"
	"DemoAuthService/internal/api/modules/profile"
	"DemoAuthService/internal/config"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type server struct {
	engine *gin.Engine
	db     *sql.Conn
	server *http.Server
}

func NewServer(db *sql.DB) *server {

	return &server{
		engine: gin.Default(),
	}
}

func (s *server) Run(conn *pgx.Conn) error {
	r := s.engine.Group("/api/v1")

	authService := auth.NewService(conn)
	authHandler := auth.NewHandler(authService)
	authHandler.RegisterRoutes(r)

	profileService := profile.NewService(conn)
	profileHandler := profile.NewHandler(profileService)
	profileHandler.RegisterRoutes(r, middlewares.Auth())

	s.engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	s.engine.GET("/err", func(ctx *gin.Context) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "test error method",
		})
	})

	return s.engine.Run(config.Default().Port)
}

// func (s *server) Stop() {
// 	s.engine.
// }
