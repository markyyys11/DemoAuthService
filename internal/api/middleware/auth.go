package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Auth(repo *sql.Conn) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
