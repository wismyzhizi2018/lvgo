package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"order/app/Common"
	"order/config"
	"os"
	"path"
	"time"
)

// Logger 获取日志的实例
func Logger() *logrus.Logger {
	logFilePath := Common.ServerInfo["storage_log_path"]
	logFileName := Common.ServerInfo["access_log_name"]
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	//写入文件
	flagExist := Common.FileIsExist(fileName)
	if !flagExist {
		if _, errs := Common.CreateFile(logFilePath, logFileName); errs != nil {
			fmt.Println("Create log file err :", errs)
		}
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//defer src.Close()
	if err != nil {
		fmt.Println("Os open log file err :", err)
	}
	//实例化
	logger := logrus.New()
	//设置输出
	logger.Out = src
	//设置日志级别
	logger.SetLevel(logrus.TraceLevel)

	//格式化时间
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

// LoggerToFiles 日志记录到文件
func LoggerToFiles() gin.HandlerFunc {
	logger := Logger()
	//color.Debug.Println(logger)
	//分割日志
	//logWriter, err := rotatelogs.New(
	//	// 分割后的文件名称
	//	fileName+".%Y%m%d.log",
	//
	//	// 生成软链，指向最新日志文件
	//	rotatelogs.WithLinkName(fileName),
	//
	//	//设置日志最大保存时间(7天)
	//	rotatelogs.WithMaxAge(7*24*time.Hour),
	//
	//	//设置日志切割时间间隔
	//	rotatelogs.WithRotationTime(24*time.Hour),
	//)
	//
	//writeMap := lfshook.WriterMap{
	//	logrus.InfoLevel:  logWriter,
	//	logrus.FatalLevel: logWriter,
	//	logrus.DebugLevel: logWriter,
	//	logrus.WarnLevel:  logWriter,
	//	logrus.ErrorLevel: logWriter,
	//	logrus.PanicLevel: logWriter,
	//}
	//lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//})
	//logger.AddHook(lfHook)
	//设置日志格式
	//logger.SetFormatter(&logrus.TextFormatter{})
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		//执行时间
		execTime := time.Since(startTime)

		//状态码
		statusCode := c.Writer.Status()

		//请求IP
		clientIP := c.ClientIP()

		//请求方式
		requestMethod := c.Request.Method

		//请求路由
		requestURI := c.Request.RequestURI

		//请求唯一ID
		requestId := c.Writer.Header().Get("X-Request-Id")

		//日志记录自定义字段
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"execute_time": execTime,
			"client_host":  clientIP,
			"http_method":  requestMethod,
			"request_url":  requestURI,
			"api":          requestURI,
			"requestId":    requestId,
		}).Info()
	}
}

var lg *zap.Logger

// InitLogger 初始化Logger
func InitLogger(cfg *config.LogConfig) (err error) {
	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return
}

//编码器(如何写入日志)
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

//指定日志将写到哪里去
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	//Filename : 日志文件的位置；
	//MaxSize ：在进行切割之前，日志文件的最大大小（以MB为单位）；
	//MaxBackups ：保留旧文件的最大个数；
	//MaxAges ：保留旧文件的最大天数；
	//Compress ：是否压缩/归档旧文件；
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

var logger *zap.Logger

func InitZapLogger(logpath string, loglevel string) {
	// 日志分割
	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径，默认 os.TempDir()
		MaxSize:    10,      // 每个日志文件保存10M，默认 100M
		MaxBackups: 30,      // 保留30个备份，默认不限
		MaxAge:     7,       // 保留7天，默认不限
		Compress:   true,    // 是否压缩，默认不压缩
	}
	write := zapcore.AddSync(&hook)
	// 设置日志级别
	// debug 可以打印出 info debug warn
	// info  级别可以打印 warn info
	// warn  只能打印 warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.RFC3339TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	core := zapcore.NewCore(
		// zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewJSONEncoder(encoderConfig),
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&write)), // 打印到控制台和文件
		write,
		level,
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段,如：添加一个服务器名称
	filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger = zap.New(core, caller, development, filed)
	logger.Info("DefaultLogger init success")
}

// ZapLoggerToFiles 日志记录到文件
func ZapLoggerToFiles(cfg *config.LogConfig) gin.HandlerFunc {
	//InitLogger(cfg)
	InitZapLogger(cfg.Filename, cfg.Level)
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		//执行时间
		execTime := time.Since(startTime)

		//状态码
		statusCode := c.Writer.Status()

		//请求IP
		clientIP := c.ClientIP()

		//请求方式
		requestMethod := c.Request.Method

		//请求路由
		requestURI := c.Request.RequestURI

		//请求唯一ID
		requestId := c.Writer.Header().Get("X-Request-Id")

		path := c.Request.URL.Path

		//日志记录自定义字段
		logger.Info(path,
			zap.Int("status_code", statusCode),
			zap.Duration("execute_time", execTime),
			zap.String("client_host", clientIP),
			zap.String("http_method", requestMethod),
			zap.String("request_url", requestURI),
			zap.String("api", requestURI),
			zap.String("requestId", requestId),
		)
	}
}
