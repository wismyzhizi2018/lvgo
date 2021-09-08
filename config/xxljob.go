package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

var xxl XXLJobConfig

type XXLJobConfig struct {
	Enabled      bool
	ServerAddr   string
	AccessToken  string
	ExecutorIp   string
	ExecutorPort interface{}
	RegistryKey  string
}

// GetXXLJobConfig 从 viper 中解析配置信息
func GetXXLJobConfig() XXLJobConfig {
	enabled := viper.GetString("JAEGER_ENABLED")
	addr := viper.GetString("XXLJOB_SERVERADDR")
	token := viper.GetString("XXLJOB_ACCESSTOKEN")
	key := viper.GetString("XXLJOB_REGISTRYKEY")
	ip := viper.GetString("XXLJOB_EXECUTORIP")
	port := viper.GetString("XXLJOB_EXECUTORPORT")
	// 默认值
	if len(enabled) == 0 {
		enabled = "false"
	}
	if enabled == "true" && key == "" {
		log.Println("XXLJOB_REGISTRYKEY 不能为空")
		os.Exit(200)
	}
	xxl.Enabled, _ = strconv.ParseBool(enabled)
	xxl.ServerAddr = addr
	xxl.AccessToken = token
	xxl.ExecutorIp = ip
	xxl.ExecutorPort = port
	xxl.RegistryKey = key
	return xxl
}
