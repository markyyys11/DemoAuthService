package profile

import (
	"DemoAuthService/internal/api/module"
	"DemoAuthService/internal/logger"

	"github.com/gin-gonic/gin"
)

type handler struct {
	ser         Service
	middlewares []gin.HandlerFunc
}

func NewHandler(s Service) module.Handler {
	return &handler{
		ser: s,
	}
}

func (h *handler) RegisterRoutes(r *gin.RouterGroup) {
	logger.Hint("Auth handler reg")
	rg := r.Group("/profile")
	{
		rg.Use(h.middlewares...)
		rg.POST("/", h.getUserInfo)
	}
}

func (h *handler) getUserInfo(ctx *gin.Context) {
	h.ser.GetUserInfo(ctx)
}
