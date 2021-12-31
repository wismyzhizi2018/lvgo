package driver

import (
	"database/sql"
	"fmt"
	"order/config"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gookit/color"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db    *gorm.DB
	dbMap map[string]*gorm.DB
)

// var MysqlDbErr error

type MysqlService interface {
	InitConnection() (map[string]*gorm.DB, error)
	Connection(conns string) (*gorm.DB, error)
	CloseConnection()
}

type mysqlService struct {
	conn       *gorm.DB
	baseConfig map[string]*config.DBSQLConf
}

func NewService(dbConfig map[string]*config.DBSQLConf) MysqlService {
	return &mysqlService{
		conn:       db,
		baseConfig: dbConfig,
	}
}

func (s *mysqlService) Connection(conns string) (*gorm.DB, error) {
	if dbSession, ok := dbMap[conns]; ok {
		return dbSession, nil
	} else {
		for connections, dbConfig := range s.baseConfig {
			if dbConfig.Driver == "mysql" && connections == conns {
				if dbSessions, err := mysqlConnect(connections, dbConfig); err == nil {
					dbMap[connections] = dbSessions
					return dbSessions, nil
				}
			}
		}
	}
	return nil, nil
}

func (s *mysqlService) CloseConnection() {
	if dbMap != nil {
		//for _, dbSession := range dbMap {
		//	dbSession.Close()
		//}
	}
}

func mysqlConnect(connections string, dbConfig *config.DBSQLConf) (*gorm.DB, error) {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
		dbConfig.Charset,
		dbConfig.Timeout,
	)
	MysqlDb, err := sql.Open("mysql", dbDSN)
	if err != nil {
		color.Danger.Println(connections+"database data source name error", err.Error())
		panic(connections + "database data source name error: " + err.Error())
	}
	// max open connections 设置打开数据库连接的最大数量
	dbMaxOpenConns, _ := strconv.Atoi(dbConfig.DbMaxOpenConns)
	MysqlDb.SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections 设置空闲连接池中连接的最大数量
	dbMaxIdleConns, _ := strconv.Atoi(dbConfig.DbMaxIdleConns)
	MysqlDb.SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <=0 will forever 设置了连接可复用的最大时间
	dbMaxLifetimeConns, _ := strconv.Atoi(dbConfig.DbMaxLifetimeConns)
	MysqlDb.SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	if err = MysqlDb.Ping(); nil != err {
		color.Danger.Println(connections+"MySQL数据库连接失败。。。", err.Error())
		os.Exit(200)
	}
	// 启用gorm
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: MysqlDb,
	}), &gorm.Config{})
	if err != nil {
		color.Danger.Println(connections+">>>GORM数据库连接失败。。。", err)
		os.Exit(200)
	} else {
		color.Info.Println(connections + "MySQL已连接GORM已连接 >>> ")
	}
	return gormDB, nil
}

func (s *mysqlService) InitConnection() (map[string]*gorm.DB, error) {
	dbMaps := make(map[string]*gorm.DB)
	for connections, dbConfig := range s.baseConfig {
		if dbConfig.Driver == "mysql" {
			if dbSession, err := mysqlConnect(connections, dbConfig); err == nil {
				if connections == "mysql" {
					db = dbSession
				}
				dbMaps[connections] = dbSession
			}
		}
	}
	dbMap = dbMaps
	return dbMaps, nil
}

//func init() {
//	log.Println("尝试连接MySQL服务...")
//
//	// get db config
//	dbConfig := config.GetMySQLConfig()
//
//	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s",
//		dbConfig["DB_USER"],
//		dbConfig["DB_PWD"],
//		dbConfig["DB_HOST"],
//		dbConfig["DB_PORT"],
//		dbConfig["DB_NAME"],
//		dbConfig["DB_CHARSET"],
//		dbConfig["DB_TIMEOUT"],
//	)
//
//	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
//
//	if MysqlDbErr != nil {
//		panic("database data source name error: " + MysqlDbErr.Error())
//	}
//
//	// max open connections
//	dbMaxOpenConns, _ := strconv.Atoi(dbConfig["DB_MAX_OPEN_CONNS"])
//	MysqlDb.SetMaxOpenConns(dbMaxOpenConns)
//
//	// max idle connections
//	dbMaxIdleConns, _ := strconv.Atoi(dbConfig["DB_MAX_IDLE_CONNS"])
//	MysqlDb.SetMaxIdleConns(dbMaxIdleConns)
//
//	// max lifetime of connection if <=0 will forever
//	dbMaxLifetimeConns, _ := strconv.Atoi(dbConfig["DB_MAX_LIFETIME_CONNS"])
//	MysqlDb.SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))
//
//	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
//		log.Println("MySQL数据库连接失败。。。", MysqlDbErr.Error())
//		//os.Exit(200)
//	} else {
//		log.Println("MySQL已连接 >>> ")
//	}
//}
