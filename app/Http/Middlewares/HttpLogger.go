package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"order/app/Common"
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

		//结束时间
		endTime := time.Now()

		//执行时间
		execTime := endTime.Sub(startTime)

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
