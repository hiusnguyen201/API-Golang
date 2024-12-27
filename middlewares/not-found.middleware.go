package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundMiddleware(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"status":  http.StatusNotFound,
		"message": http.StatusText(http.StatusNotFound),
	})
}