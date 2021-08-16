package Common

// 项目函数运行所依赖的参数配置

import (
	"order/config"
)

// Config 代码参数配置
var Config = map[string]interface{}{
	"log_debug": config.GetFrameworkConfig()["log_debug"],
	"api":       config.GetFrameworkConfig()["api"],
}

// ServerInfo 框架和服务参数
var ServerInfo = map[string]string{
	"timezone":         config.GetFrameworkConfig()["timezone"], // 时区
	"gl_version":       config.GetFrameworkConfig()["gl_version"],
	"go_version":       config.GetFrameworkConfig()["go_version"],
	"go_root":          config.GetFrameworkConfig()["go_root"],
	"framework_path":   config.GetFrameworkConfig()["framework_path"],
	"storage_path":     config.GetFrameworkConfig()["storage_path"],          // 文件存储在服务器的地址
	"storage_log_path": config.GetFrameworkConfig()["storage_path"] + "logs", // 日志文件文存储在服务器的地址
	"access_log_name":  "access_go_laravel.log",                              // 访问的文件名称
	"error_log_name":   "error_go_laravel.log",                               // 错误的文件名称
	"slow_log_name":    "slow_go_laravel.log",                                // 慢运行文件名称
}

// Page 分页参数
var Page = map[string]int{
	"limit":     20,   // 每页多少条数据
	"max_page":  1000, // 最大查到多少页
	"min_limit": 2,    // 最少一次查多少条数据
	"max_limit": 100,  // 最多一次查多少条数据
}
