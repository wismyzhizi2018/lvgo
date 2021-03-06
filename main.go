package main

import (
	_ "embed"
	"order/bootstrap"
	"order/config"

	"github.com/gin-gonic/gin"
	"github.com/namsral/flag"
)

var HttpServer *gin.Engine

//go:embed .env
var BytesContent []byte

type Application interface {
	App()
}

func bootstraps(app Application) {
	app.App()
}

//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
var endpoint = flag.String("endpoint", "<point>", "nacos endpoint")

var (
	namespaceId = flag.String("namespace_id", "<namespace_id>", "nacos namespace Id")
	accessKey   = flag.String("access_key", "<access_key>", "nacos access key")
	secretKey   = flag.String("secret_key", "<secret_key>", "nacos secret key")
	dataId      = flag.String("data_id", "app.yaml", "nacos secret key")
	group       = flag.String("group", "test", "nacos secret key")
	port        = flag.Uint64("port", 8080, "nacos port")
)

func main() {
	flag.Parse()
	config.NewNacosConfig(*endpoint, *namespaceId, *accessKey, *secretKey, *dataId, *group, *port)
	// 启动服务
	var app Application = &bootstrap.Application{HttpServer: HttpServer, BytesContent: BytesContent}
	bootstraps(app)
}
