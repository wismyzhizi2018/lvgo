package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

//func InitConfig() {
//	DatabaseConfig = GetDBSQLConf()
//}

var DatabaseConfig map[string]*DBSQLConf

// DBSQLConf 结构体中的成员变量，只有首字母大写，才能在其定义的 package 以外访问。而在同一个 package 内，就不会有此限制。
// 数据库的配置
type DBSQLConf struct {
	Driver             string // 驱动名称
	Host               string // 地址
	Port               string // 端口
	Database           string // 数据库名称
	Username           string // 用户名
	Password           string // 密码
	Charset            string // 字符编码
	Timeout            string // 超时时间
	DbMaxOpenConns     string // 连接池最大连接数
	DbMaxIdleConns     string // 连接池最大空闲数
	DbMaxLifetimeConns string // 连接池链接最长生命周期
}

// GetDataBaseConfig MySQL数据库配置
func GetDataBaseConfig() map[string]*DBSQLConf {
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	name := viper.GetString("DB_DATABASE")
	user := viper.GetString("DB_USERNAME")
	// 默认值
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	if len(port) == 0 {
		port = "3306"
	}
	if len(name) == 0 || len(user) == 0 {
		log.Println("mysql默认数据库名或账户名不能为空。。。")
		os.Exit(200)
	}
	databaseConf := make(map[string]*DBSQLConf)

	// 数据库默认配置
	databaseConf["mysql"] = &DBSQLConf{
		Driver:             "mysql",
		Host:               viper.GetString("DB_HOST"),
		Port:               viper.GetString("DB_PORT"),
		Database:           viper.GetString("DB_DATABASE"),
		Username:           viper.GetString("DB_USERNAME"),
		Password:           viper.GetString("DB_PASSWORD"),
		Charset:            "utf8mb4",
		Timeout:            "12s",
		DbMaxOpenConns:     "20",
		DbMaxIdleConns:     "10",
		DbMaxLifetimeConns: "7200",
	}
	// mysql日志库
	databaseConf["mysql_order_log"] = &DBSQLConf{
		Driver:             "mysql",
		Host:               viper.GetString("DB_ORDER_LOG_HOST"),
		Port:               viper.GetString("DB_ORDER_LOG_PORT"),
		Database:           viper.GetString("DB_ORDER_LOG_DATABASE"),
		Username:           viper.GetString("DB_ORDER_LOG_USERNAME"),
		Password:           viper.GetString("DB_ORDER_LOG_PASSWORD"),
		Charset:            "utf8mb4",
		Timeout:            "12s",
		DbMaxOpenConns:     "20",
		DbMaxIdleConns:     "10",
		DbMaxLifetimeConns: "7200",
	}

	// mongo db的配置
	databaseConf["mongodb"] = &DBSQLConf{
		Driver:             "mongodb",
		Host:               viper.GetString("MONGODB_HOST"),
		Port:               viper.GetString("MONGODB_PORT"),
		Database:           viper.GetString("MONGODB_DATABASE"),
		Username:           viper.GetString("MONGODB_USERNAME"),
		Password:           viper.GetString("MONGODB_PASSWORD"),
		Charset:            "utf8mb4",
		Timeout:            "12s",
		DbMaxOpenConns:     "20",
		DbMaxIdleConns:     "10",
		DbMaxLifetimeConns: "7200",
	}

	// sqlsrv db的配置
	databaseConf["sqlsrv"] = &DBSQLConf{
		Driver:             "sqlsrv",
		Host:               viper.GetString("SQL_SERVER_HOST"),
		Port:               viper.GetString("SQL_SERVER_PORT"),
		Database:           viper.GetString("SQL_SERVER_DATABASE"),
		Username:           viper.GetString("SQL_SERVER_USERNAME"),
		Password:           viper.GetString("SQL_SERVER_PASSWORD"),
		Charset:            "utf8mb4",
		Timeout:            "12s",
		DbMaxOpenConns:     "20",
		DbMaxIdleConns:     "10",
		DbMaxLifetimeConns: "7200",
	}

	// pgsql db的配置
	databaseConf["pgsql"] = &DBSQLConf{
		Driver:             "sqlsrv",
		Host:               viper.GetString("DB_HOST"),
		Port:               viper.GetString("DB_PORT"),
		Database:           viper.GetString("DB_DATABASE"),
		Username:           viper.GetString("DB_USERNAME"),
		Password:           viper.GetString("DB_PASSWORD"),
		Charset:            "utf8mb4",
		Timeout:            "12s",
		DbMaxOpenConns:     "20",
		DbMaxIdleConns:     "10",
		DbMaxLifetimeConns: "7200",
	}
	// 更多..
	return databaseConf
}

func GetRedisConfig() map[string]string {
	conf := make(map[string]string)
	host := viper.GetString("REDIS_HOST")
	pwd := viper.GetString("REDIS_PASSWORD")
	port := viper.GetString("REDIS_PORT")
	addr := host + ":" + port
	conf["Addr"] = addr    // 例子：127.0.0.1:6379
	conf["Password"] = pwd // 无密码就设置为：""
	conf["DB"] = "0"       // use default DB

	return conf
}
