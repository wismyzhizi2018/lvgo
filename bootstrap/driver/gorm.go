package driver

import (
	"database/sql"
	"fmt"
	"order/config"

	"github.com/gookit/color"
	"gorm.io/driver/mysql"

	//"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	GDB  map[string]*gorm.DB
	gErr error
)

func InitGorm() {
	color.Debug.Println(">>>GORM开始接现有数据库驱动 >>> ")
	baseConfig := config.GetDataBaseConfig()
	dbMaps := make(map[string]*gorm.DB)

	for connections, dbConfig := range baseConfig {
		if dbConfig.Driver == "mysql" {
			dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s",
				dbConfig.Username,
				dbConfig.Password,
				dbConfig.Host,
				dbConfig.Port,
				dbConfig.Database,
				dbConfig.Charset,
				dbConfig.Timeout,
			)
			// 连接现有MySQL
			sqlDB, sErr := sql.Open("mysql", dbDSN)
			if sErr != nil {
				color.Danger.Println(connections+">>>GORM现有数据库连接失败，GORM功能将不可用。。。", sErr)
				// os.Exit(200)
			} else {
				color.Info.Println(connections + ">>>尝试连接GORM... ")
			}
			var GDBS *gorm.DB
			GDBS, gErr = gorm.Open(mysql.New(mysql.Config{
				Conn: sqlDB,
			}), &gorm.Config{})
			dbMaps[connections] = GDBS

			// fmt.Println(GDB)

			if gErr != nil {
				color.Danger.Println(connections+">>>GORM数据库连接失败。。。", gErr)
				// os.Exit(200)
			} else {
				color.Info.Println(connections + ">>>GORM已连接现有数据库驱动 >>> ")
			}
		}
	}
	GDB = dbMaps
}
