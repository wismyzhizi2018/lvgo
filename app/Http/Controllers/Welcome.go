package Controllers

import (
	"math"
	"order/app/Common"
	//"ginvel.com/app/Kit"
	"order/config"

	"github.com/gin-gonic/gin"
	util "order/app/Helpers"
)

// Welcome
// @title 默认路由
// @description 默认路由不用设置其他功能
// @Router / [GET]
func Welcome(ctx *gin.Context) {
	// 读取超全局变量即可
	var _cpuPercent interface{} = Common.GetGlobalData("cpu_percent")
	var __cpuPercent float64 = _cpuPercent.(float64)
	var cpuPercent int64 = int64(math.Floor(__cpuPercent))
	// fmt.Println(Kit.RDB.Set(ctx, "go_database_jwt_cache:15999640681", "123", time.Second*15))
	// fmt.Println(Kit.RDB.Set(ctx, "15999640682", "456", time.Second*15))
	// fmt.Println(Kit.RDB.Set(ctx, "15999640683", "123", time.Second*60))
	// fmt.Println(Kit.RDB.Get(ctx, "15999640682"))
	// fmt.Println(Kit.RDB.Get(ctx, "15999640681"))
	// fmt.Println(Kit.RDB.Get(ctx, "15999640683"))

	back := map[string]interface{}{
		"state": 1,
		"msg":   "接口请求成功，进入默认路由",
		"content": map[string]interface{}{
			"gl_version":  config.GetFrameworkConfig()["gl_version"],
			"go_version":  config.GetFrameworkConfig()["go_version"],
			"timezone":    config.GetFrameworkConfig()["timezone"],
			"time":        Common.GetTimeDate("Ymd.His.ms.ns"),
			"cpu_percent": cpuPercent,
		},
	}
	util.SuccessJson("请求成功", back)(ctx)
	// ctx.JSONP(http.StatusOK, back)
}

// Api
// @title CorsApi示例
// @description 测试api
// @Router / [GET]
func Api(ctx *gin.Context) {
	name := ctx.Query("name")

	if len(name) == 0 {
		name = "name为空"
		util.FailJson(1, name, gin.H{}, gin.H{})(ctx)
		return
	}
	_id := ctx.Query("id")
	id := Common.StringToInt(_id)

	content := map[string]interface{}{
		"name": name,
		"id":   id,
	}

	back := map[string]interface{}{
		"content": content,
	}

	util.SuccessJson("接口请求成功，进入CorsApi示例", back)(ctx)
}
