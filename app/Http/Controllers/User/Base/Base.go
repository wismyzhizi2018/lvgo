package Base

import (
	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	user, flag := ctx.Get("authedUserObj")
	//fmt.Println(user)
	//fmt.Println(Order.GetOrder(orderCode))
	if flag != true {
		ctx.JSON(200, gin.H{"error": "user not exists"})
	}
	// 接口返回
	back := gin.H{
		"userInfo": user,
	}
	ctx.JSON(200, back)
}
