package bootstrap

import (
	"github.com/arl/statsviz"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"log"
	"net/http"
	"order/app/Http/Middlewares"
	"order/app/Http/Models/Kit"
	"order/app/Http/Request"
	"order/bootstrap/driver"
	"order/config"
	"order/routes"
	"os"
)

type Application struct {
	HttpServer   *gin.Engine
	BytesContent []byte
}

// App 配置并启动http服务
// 项目访问入口
func (app *Application) App() {

	// 必选初始化
	HttpServer := app.HttpServer
	//加载配置信息
	config.InitEmbedData(app.BytesContent)

	//获取配置信息
	serverConfig := config.GetServerConfig()
	frameworkConfig := config.GetFrameworkConfig()
	Request.InitTrans(frameworkConfig["locale"])

	//加载mysql链接
	baseInfo := config.GetDataBaseConfig()
	driver.NewService(baseInfo).InitConnection()
	driver.InitGorm()
	driver.InitRedis()
	driver.InitMongo()

	DB, _ := driver.NewService(baseInfo).GetMYSQLConnection("mysql")
	Kit.Db = DB
	Kit.RDB = driver.RedisDb
	Kit.MDB = driver.MongodbClient
	Kit.DB = driver.GDB["mysql"]
	defer driver.NewService(baseInfo).CloseConnection()
	// 服务停止时清理数据库链接

	//svc.InitConnection()
	// Gin服务
	HttpServer = gin.New()

	if serverConfig["ENV"] == "debug" {
		pprof.Register(HttpServer)
		go func() {
			if err := statsviz.RegisterDefault(); err != nil {
				color.Danger.Println("开启SV监控失败", err.Error())
			}
			color.Debug.Println(http.ListenAndServe("localhost:8080", nil))
		}()
	}

	//启用链路追踪中间件
	HttpServer.Use(Middlewares.RequestId)

	jaegerConfig := config.ParseConfig()
	if jaegerConfig.Enabled {
		//启用链路追踪中间件
		color.Debug.Println("启用链路追踪中间件 ")
		HttpServer.Use(Middlewares.Jaeger())
	}

	//启用日志中间件
	HttpServer.Use(Middlewares.LoggerToFiles())

	// 捕捉接口运行耗时（必须排第一）
	HttpServer.Use(Middlewares.StatLatency)

	// 设置全局ctx参数（必须排第二）
	HttpServer.Use(Middlewares.AppData)

	// 拦截应用500报错，使之可视化
	HttpServer.Use(Middlewares.AppError500)

	// Gin运行时：release、debug、test
	//	log.Println(serverConfig["ENV"])
	gin.SetMode(serverConfig["ENV"])

	// 注册必要路由，处理默认路由、静态文件路由、404路由等
	routes.RouteMust(HttpServer)

	// 注册其他路由，可以自定义

	if serverConfig["ENV"] == "debug" {
		color.Debug.Println("开启  pprof>>> ")
		pprof.RouteRegister(routes.RouterRegister(HttpServer))
	} else {
		routes.RouterRegister(HttpServer)
	}

	//Router.Api(HttpServer) // 面向Api
	//Router.Web(HttpServer) // 面向模版输出

	// 初始化定时器（立即运行定时器）
	//Task.TimeInterval(0, 0, "0")

	// 实际访问网址和端口
	_host := "127.0.0.1:" + serverConfig["PORT"]              // 测试访问IP
	host := serverConfig["HOST"] + ":" + serverConfig["PORT"] // Docker访问IP

	glVersion := frameworkConfig["gl_version"]

	shellMessage := " \n " +
		"访问地址示例：" + host + " >>> \n " +
		"gl_version：" + glVersion + " \n " +
		"1) 默认接口（JSON输出）：http://" + _host + " \n " +
		"2) 测试接口（JSON输出）：http://" + _host + "/api?name=gl&id=2021 \n " +
		"3) 静态文件输出（文件）：http://" + _host + "/favicon.ico \n " +
		"4) 查看WebSocket连接：ws://" + _host + "/api/example/socket/ping1 \n " +
		""
	if serverConfig["ENV"] == "debug" {
		shellMessage = shellMessage + "5) SV系统集成可视化实时运行时统计：http://localhost:8080/debug/statsviz/ \n "
	}
	// 终端提示
	color.Debug.Println(shellMessage)

	err := HttpServer.Run(host)
	if err != nil {
		log.Println("http服务遇到错误，运行中断，error：", err.Error())
		log.Println("提示：注意端口被占时应该首先更改对外暴露的端口，而不是微服务的端口。")
		os.Exit(200)
	}

	return
}
