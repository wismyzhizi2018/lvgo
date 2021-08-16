package Middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId(ctx *gin.Context) {
	if ctx.Request.Header.Get("X-Request-Id") == "" {
		requestId, _ := getUUID()
		ctx.Set("X-Request-Id", requestId)
		ctx.Writer.Header().Set("X-Request-Id", requestId)
	}
	ctx.Next()
}

func getUUID() (string, error) {
	u2 := uuid.NewV4()
	return u2.String(), nil
}
