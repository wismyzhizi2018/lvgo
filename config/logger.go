package config

import (
	"os"
	"path"
)

var logConfig LogConfig

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

// NewLogConfig 从 viper 中解析配置信息
func NewLogConfig() LogConfig {
	logConfig.Level = "info"
	logConfig.Filename = "access_log"
	logConfig.MaxSize = 1024
	logConfig.MaxAge = 1024
	logConfig.MaxBackups = 1024
	return logConfig
}

func GetLogConfig() *LogConfig {
	logFilePath, _ := os.Getwd()
	logFileName := "/storage/logs/access_go_laravel.log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	logConfig.Level = "info"
	logConfig.Filename = fileName
	logConfig.MaxSize = 10
	logConfig.MaxAge = 5
	logConfig.MaxBackups = 30
	return &logConfig
}
