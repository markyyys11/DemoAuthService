package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer ctx.Next()

		token := strings.TrimSpace(ctx.Request.Header.Get("Authorization"))

		if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "No Authorization header",
			})
			ctx.Abort()
			return
		}

		if !strings.Contains(token, "Bearer") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Error parse Authorization header",
			})
			ctx.Abort()
			return
		}
	}

}
