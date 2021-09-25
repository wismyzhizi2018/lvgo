package bootstrap

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xxl-job/xxl-job-executor-go/example/task"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"order/app/Libs/Consul"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/satori/go.uuid"

	"github.com/arl/statsviz"
	"github.com/gin-contrib/pprof"
	"github.com/gin-middleware/xxl-job-executor"
	"github.com/gookit/color"
	"github.com/xxl-job/xxl-job-executor-go"
	"order/app/Http/Middlewares"
	"order/app/Http/Models/Kit"
	"order/app/Http/Request"
	"order/bootstrap/driver"
	"order/config"
	"order/routes"
)

type Application struct {
	HttpServer   *gin.Engine
	Logger       *zap.Logger
	BytesContent []byte
}

// App 配置并启动http服务
// 项目访问入口
func (app *Application) App() {

	// 必选初始化
	HttpServer := app.HttpServer
	//加载配置信息
	config.InitEmbedData(app.BytesContent)
	//config.InitNACOS()
	mainDirectory, _ := os.Getwd()
	mainDirectory = mainDirectory + "/"
	bytes, err := ioutil.ReadFile(mainDirectory + "config/banner.txt")
	if err != nil {
		nameField := zap.String("name", "banner s ")
		app.Logger.Fatal("banner error", nameField, zap.Error(err))
		return
	}
	color.Info.Printf("%v\n", string(bytes))
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
	//HttpServer.Use(Middlewares.LoggerToFiles())
	HttpServer.Use(Middlewares.ZapLoggerToFiles(config.GetLogConfig()))

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

	//xxl-job
	xxlConfig := config.GetXXLJobConfig()
	//初始化执行器
	if xxlConfig.Enabled {
		exec := xxl.NewExecutor(
			xxl.ServerAddr(xxlConfig.ServerAddr),
			xxl.AccessToken(xxlConfig.AccessToken), //请求令牌(默认为空)
			xxl.ExecutorIp(xxlConfig.ExecutorIp),   //可自动获取
			xxl.ExecutorPort(serverConfig["PORT"]), //默认9999（此处要与gin服务启动port必需一至）
			xxl.RegistryKey(xxlConfig.RegistryKey), //执行器名称
			//xxl.SetLogger(&logger{}), //执行器名称
		)
		exec.Init()
		defer exec.Stop()
		//设置日志查看handler
		//exec.LogHandler(func(req *xxl.LogReq) *xxl.LogRes {
		//	return &xxl.LogRes{Code: 200, Msg: "", Content: xxl.LogResContent{
		//		FromLineNum: req.FromLineNum,
		//		ToLineNum:   2,
		//		LogContent:  "这个是自定义日志handler",
		//		IsEnd:       true,
		//	}}
		//})
		xxl_job_executor_gin.XxlJobMux(HttpServer, exec)
		//注册任务handler
		exec.RegTask("task.test", task.Test)
		exec.RegTask("task.test2", task.Test2)
		exec.RegTask("task.panic", task.Panic)
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

	srv := &http.Server{
		Addr:    ":" + serverConfig["PORT"],
		Handler: HttpServer,
	}

	//后台启动一个goroutine来启动服务
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			color.Debug.Printf("listen: %s\n", err)
			os.Exit(1)
		}
	}()

	//使用注册中心上报id和服务器状态
	consulConfig := config.GetConsulConfig()

	host = consulConfig["HOST"].(string)
	port, _ := strconv.Atoi(consulConfig["PORT"].(string))
	serverPort, _ := strconv.Atoi(serverConfig["PORT"])
	registerClient := Consul.NewRegistryClient(host, port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	if consulConfig["Enabled"].(bool) {
		err := registerClient.Register("192.168.0.49", serverPort, "lvgo-services", []string{"php", "go", "laravel", "gin"}, serviceId)
		if err != nil {
			color.Danger.Println("consul:", err)
			os.Exit(1)
		}
	}
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	color.Info.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		color.Danger.Println("Server Shutdown:", err)
		os.Exit(1)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		color.Info.Println("timeout of 5 seconds.")
	}
	if consulConfig["Enabled"].(bool) {
		if errs := registerClient.DeRegister(serviceId); errs != nil {
			color.Danger.Printf("注销console ,%s", errs)
		}
	}
	color.Info.Printf("注销console ,%s", "success")
	color.Info.Println("exit Server ...")
	//err := HttpServer.Run(host)
	//if err != nil {
	//	log.Println("http服务遇到错误，运行中断，error：", err.Error())
	//	log.Println("提示：注意端口被占时应该首先更改对外暴露的端口，而不是微服务的端口。")
	//	os.Exit(200)
	//}
	return
}

//type logger struct{}
//
//func (l *logger) Info(format string, a ...interface{}) {
//	fmt.Println(fmt.Sprintf("自定义日志 - "+format, a...))
//}
//
//func (l *logger) Error(format string, a ...interface{}) {
//	log.Println(fmt.Sprintf("自定义日志 - "+format, a...))
//}
