package config

import (
	"bytes"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func LoadInit() {
	arg1 := strings.ToLower(os.Args[0])
	name := filepath.Base(arg1)
	flagMain := strings.Index(name, "main") == 0 && strings.Index(arg1, "go-build") < 0
	flagServer := strings.Index(name, "webserver") == 0 && strings.Index(arg1, "go-build") < 0
	if flagMain || flagServer {
		//从二进制中InitBinData加载配置文件
		log.Println("运行加载InitBinData下的env配置项文件 >>> ")
		//InitBinData()
	} else {
		log.Println("运行加载env配置项文件 >>> ")
		//	InitEnv()
	}
}

func InitEnv() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() //根据上面配置加载文件
	if err != nil {
		log.Println("env decode Error = ", err.Error(), "运行中断")
		os.Exit(200)
	}

}

func InitEmbedData(bytesContent []byte) {
	//设置要读取的文件类型
	color.Debug.Println("使用go:embed加载配置文件")
	viper.SetConfigType("env")
	//读取
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
	color.Danger.Println(locale)
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
	color.Danger.Println(locale)

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
