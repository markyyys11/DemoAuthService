package auth

import (
	"DemoAuthService/internal/api/modules"
	"DemoAuthService/internal/logger"

	"github.com/gin-gonic/gin"
)

type handler struct {
	ser Service
}

func NewHandler(s Service) modules.Handler {
	return &handler{
		ser: s,
	}
}

func (h *handler) RegisterRoutes(r *gin.RouterGroup, mw ...gin.HandlerFunc) {
	logger.Hint("Auth handler reg")
	rg := r.Group("/auth")
	{
		rg.Use(mw...)
		rg.GET("/register", h.register)
		rg.POST("/login", h.login)
		rg.GET("/logout", h.logout)
		rg.GET("/refresh", h.refresh)
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

func (h *handler) refresh(ctx *gin.Context) {
	h.ser.Refresh(ctx)
}
