package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/namsral/flag"
	"log"
	"order/bootstrap"
	"order/config"
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
var endpoint = flag.String("endpoint", "<point>", "nacos endpoint")
var namespaceId = flag.String("namespace_id", "<namespace_id>", "nacos namespace Id")
var accessKey = flag.String("access_key", "<access_key>", "nacos access key")
var secretKey = flag.String("secret_key", "<secret_key>", "nacos secret key")
var dataId = flag.String("data_id", "app.yaml", "nacos secret key")
var group = flag.String("group", "test", "nacos secret key")
var port = flag.Uint64("port", 8080, "nacos port")

func main() {
	flag.Parse()
	config.NewNacosConfig(*endpoint, *namespaceId, *accessKey, *secretKey, *dataId, *group, *port)
	// 启动服务
	var app Application = &bootstrap.Application{HttpServer: HttpServer, BytesContent: BytesContent}
	bootstraps(app)
}
