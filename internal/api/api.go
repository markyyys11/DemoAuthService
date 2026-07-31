package api

import (
	"DemoAuthService/internal/config"
	"DemoAuthService/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server struct {
	eng *gin.Engine
}

func NewServer() *server {
	return &server{
		eng: gin.Default(),
	}
}

func (s *server) RegisterHandlers(handlers ...handlers.Handler) {
	r := s.eng.Group("/api")
	for _, h := range handlers {
		h.RegisterRoutes(r)
	}
}

func (s *server) Run() error {
	// Define a simple GET endpoint
	s.eng.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return s.eng.Run(config.Default().Addr)
}
