package Helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessJson(message string, data interface{}) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  message,
			"data": data,
		})
	}
}

func FailJson(code int, message string, data, input interface{}) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  code,
			"msg":   message,
			"data":  data,
			"input": input,
		})
	}
}
