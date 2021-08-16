package Kit

// 定义全局数据库主语法，这样就可以避免空间命名混乱造成的全局数据库主语法定义位置多的问题。
// 引用如：Kit.Kit.Table("gl_user").xxx

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var Db *sql.DB        // 连接mysql扩展
var DB *gorm.DB       // 连接gorm扩展
var RDB *redis.Client // 连接Redis扩展
var MDB *mongo.Client // 连接MongoDb扩展

//driver.NewService(config.DatabaseConfig).GetMYSQLConnection("mysql")
//定义一个接口

//获取获取mysql链接
//获取reid连接
//获取gorm链接
