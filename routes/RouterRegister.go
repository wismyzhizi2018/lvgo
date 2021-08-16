package routes

import (
	"fmt"
	"order/routes/RouterGroup"

	"github.com/gin-gonic/gin"
)

// RouterRegister 灵活的注册路由文件
func RouterRegister(app *gin.Engine) (group *gin.RouterGroup) {
	fmt.Println("运行自定义注册路由文件 >>> ")
	// 示例
	adminGroup := RouterGroup.ExampleApi(app) // 面向Api

	// 其他
	// RouterGroup.ApiGen1(app)
	// RouterGroup.ApiGen3(app)
	//log.Println("log测试 \n ")

	return adminGroup
}
