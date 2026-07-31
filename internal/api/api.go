package api

import (
	"DemoAuthService/internal/api/module"
	"DemoAuthService/internal/config"
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

func (s *server) RegisterHandlers(handlers ...module.Handler) {
	r := s.eng.Group("/api/v1")
	for _, h := range handlers {
		h.RegisterRoutes(r)
	}
}

func (s *server) Run() error {
	s.eng.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return s.eng.Run(config.Default().Addr)
}
