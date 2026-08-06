package profile

import (
	"DemoAuthService/internal/api/modules"
	"net/http"

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
	rg := r.Group("/profile")
	{
		rg.Use(mw...)
		rg.POST("/", h.getUserInfo)
	}
}

func (h *handler) getUserInfo(ctx *gin.Context) {
	h.ser.GetUserInfo(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"id":   1,
		"name": "mark",
	})
}
