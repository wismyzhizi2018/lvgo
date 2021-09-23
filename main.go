package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"order/bootstrap"
	"os"
)

var HttpServer *gin.Engine

//go:embed .env.example
var BytesContent []byte

type Application interface {
	App()
}

func init() {
	// 日志写入文件
	// 已 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
	logFile, err := os.OpenFile(`./order_system.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// 设置存储位置
	log.SetOutput(logFile)
}

func bootstraps(app Application) {
	app.App()
}

//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
func main() {
	// 启动服务
	var app Application = &bootstrap.Application{HttpServer: HttpServer, BytesContent: BytesContent}
	bootstraps(app)
}
