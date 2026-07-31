package auth

import (
	"DemoAuthService/internal/handlers"
	"DemoAuthService/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	s Service
}

func NewHandler(s Service) handlers.Handler {
	return &handler{
		s: s,
	}
}

func (h *handler) RegisterRoutes(r *gin.RouterGroup) {
	logger.Hint("Auth handler reg")
	rg := r.Group("/auth")
	{
		rg.POST("/login", h.login)
	}
}

func (h *handler) login(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
