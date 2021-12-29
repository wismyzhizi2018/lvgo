package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/viper"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func LoadInit() {
	arg1 := strings.ToLower(os.Args[0])
	name := filepath.Base(arg1)
	flagMain := strings.Index(name, "main") == 0 && strings.Index(arg1, "go-build") < 0
	flagServer := strings.Index(name, "webserver") == 0 && strings.Index(arg1, "go-build") < 0
	if flagMain || flagServer {
		// 从二进制中InitBinData加载配置文件
		log.Println("运行加载InitBinData下的env配置项文件 >>> ")
		// InitBinData()
	} else {
		log.Println("运行加载env配置项文件 >>> ")
		//	InitEnv()
	}
}

func InitEnv() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // 根据上面配置加载文件
	if err != nil {
		log.Println("env decode Error = ", err.Error(), "运行中断")
		os.Exit(200)
	}
}

// InitNACOS
// @title  从配置中心读取配置文件
// @description  从配置中心读取配置文件
func InitNACOS() {
	mainDirectory, _ := os.Getwd()
	logFilePath := mainDirectory + "/tmp/nacos/log/"
	logFileName := "nacos.log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	_, err := os.Stat(fileName)
	if !(err == nil || os.IsExist(err)) {
		var err error
		// 目录不存在则创建
		if _, err = os.Stat(logFilePath); err != nil {
			if err = os.MkdirAll(logFilePath, 0777); err != nil { // 这里如果是0711权限 可能会导致其它线程，读取文件夹内内容出错
				color.Danger.Println("Create log dir err :", err)
			}
		}
		// 创建文件
		if _, err = os.Create(fileName); err != nil {
			color.Danger.Println("Create log file err :", err)
		}
	}
	nacosConf := GetNacosConfig()

	sc := []constant.ServerConfig{
		{
			IpAddr: nacosConf.Endpoint,
			Port:   nacosConf.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         nacosConf.NamespaceId, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5 * 1000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		AccessKey:           nacosConf.AccessKey,
		SecretKey:           nacosConf.SecretKey,
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
		ListenInterval:      30 * 1000,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		color.Danger.Println("nacos read Error = ", err.Error(), "运行中断")
		fmt.Println(err.Error())
		os.Exit(200)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosConf.DataId,
		Group:  nacosConf.Group,
	})
	if err != nil {
		color.Danger.Println("env read Error = ", err.Error(), "运行中断")
		fmt.Println(err.Error())
		os.Exit(200)
	}
	color.Info.Println(content) // 字符串 - yaml
	color.Debug.Println("使用NACOS加载配置文件")
	viper.SetConfigType("env")
	// 读取
	if err := viper.ReadConfig(bytes.NewBuffer([]byte(content))); err != nil {
		color.Danger.Println("env read Error = ", err.Error(), "运行中断")
		fmt.Println(err.Error())
		os.Exit(200)
	}
}

func InitEmbedData(bytesContent []byte) {
	// 设置要读取的文件类型
	if len(bytesContent) == 0 {
		color.Error.Println("配置文件为空,请检查读取配置文件是否正确")
		os.Exit(200)
	}
	color.Debug.Println("使用go:embed加载配置文件")
	viper.SetConfigType("env")
	// 读取
	if err := viper.ReadConfig(bytes.NewBuffer(bytesContent)); err != nil {
		log.Println("env read Error = ", err.Error(), "运行中断")
		os.Exit(200)
	}
}

//func InitBinData() {
//	//从打包后的文件中读取配置
//	bytesContent, err := bindata.Asset(".env")
//	if err != nil {
//		log.Println("env decode Error = ", err.Error(), "运行中断")
//		os.Exit(200)
//	}
//	//设置要读取的文件类型
//	viper.SetConfigType("env")
//	//读取
//	if err := viper.ReadConfig(bytes.NewBuffer(bytesContent)); err != nil {
//		log.Println("env read Error = ", err.Error(), "运行中断")
//		os.Exit(200)
//	}
//}

func GetServerConfig() map[string]string {
	host := viper.GetString("APP_HOST")
	port := viper.GetString("APP_PORT")
	env := viper.GetString("APP_ENV")
	// 默认值
	// docker中运行请使用：0.0.0.0，本地测试请使用：127.0.0.1
	if host == "localhost" || len(host) == 0 {
		host = "0.0.0.0"
	}
	if len(port) == 0 {
		port = "8090"
	}
	if len(env) == 0 || env == "production" {
		env = "release"
	} else if env == "local" {
		env = "debug"
	}
	conf := make(map[string]string)
	conf["HOST"] = host // 监听地址，部署在docker中请使用：0.0.0.0。建议不要用127.0.0.1或localhost
	conf["PORT"] = port // 监听端口
	conf["ENV"] = env   // 环境模式 release/debug/test
	return conf
}

// GetFrameworkConfig 框架参数配置
func GetFrameworkConfig() map[string]string {
	timezone := viper.GetString("APP_TIMEZONE")
	locale := viper.GetString("LOCALE")
	//	color.Danger.Println(locale)
	// main.go文件的绝对路径
	mainDirectory, _ := os.Getwd()
	mainDirectory = mainDirectory + "/"
	if len(timezone) == 0 {
		timezone = "Asia/Shanghai"
	}
	var logDebug string
	env := viper.GetString("APP_ENV")
	if len(env) == 0 || env == "production" {
		logDebug = "false"
	} else if env == "local" {
		logDebug = "true"
	} else {
		logDebug = "false"
	}
	if len(locale) == 0 {
		locale = "en"
	}
	// color.Danger.Println(locale)

	conf := make(map[string]string)
	conf["timezone"] = timezone // 时区
	conf["gl_version"] = "gl_version.default.0.0.0"
	conf["go_version"] = runtime.Version() // go版本
	conf["go_root"] = runtime.GOROOT()
	conf["framework_path"] = mainDirectory            // 默认使用框架根目录
	conf["storage_path"] = mainDirectory + "storage/" // 文件存储文件夹
	conf["log_debug"] = logDebug
	conf["locale"] = locale
	return conf
}

// GetViewConfig html模版视图路径配置
func GetViewConfig() map[string]string {
	pattern := viper.GetString("ViewPattern")
	static := viper.GetString("ViewStatic")
	// 默认值
	if len(pattern) == 0 {
		pattern = "views/html/**/**/*"
	}
	if len(static) == 0 {
		static = "views/static/"
	}
	conf := make(map[string]string)
	// html模板文件路径。**代表文件夹，*代表文件。*结尾。
	conf["View_Pattern"] = pattern
	// 多静态文件的主文件夹。/结尾。
	conf["View_Static"] = static
	return conf
}

// GetConsulConfig 框架参数配置
func GetConsulConfig() map[string]interface{} {
	host := viper.GetString("CONSUL_HOST")
	port := viper.GetString("CONSUL_POST")
	conf := make(map[string]interface{})
	conf["HOST"] = host // 监听地址，部署在docker中请使用：0.0.0.0。建议不要用127.0.0.1或localhost
	conf["PORT"] = port // 监听端口
	conf["Enabled"], _ = strconv.ParseBool(viper.GetString("CONSUL_ENABLED"))
	return conf
}
