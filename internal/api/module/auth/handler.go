package auth

import (
	"DemoAuthService/internal/api/module"
	"DemoAuthService/internal/logger"

	"github.com/gin-gonic/gin"
)

type handler struct {
	ser Service
}

func NewHandler(s Service) module.Handler {
	return &handler{
		ser: s,
	}
}

func (h *handler) RegisterRoutes(r *gin.RouterGroup) {
	logger.Hint("Auth handler reg")
	rg := r.Group("/auth")
	{
		rg.Use()
		rg.POST("/login", h.login)
	}
}

func (h *handler) register(ctx *gin.Context) {
	h.ser.Register(ctx)
}

func (h *handler) login(ctx *gin.Context) {
	h.ser.Login(ctx)
}

func (h *handler) logout(ctx *gin.Context) {
	h.ser.Logout(ctx)
}
