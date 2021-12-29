package driver

import (
	"database/sql"
	"fmt"
	"order/config"
	"strconv"
	"time"

	"github.com/gookit/color"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db    *sql.DB
	dbMap map[string]*sql.DB
)

// var MysqlDbErr error

type MysqlService interface {
	InitConnection() (map[string]*sql.DB, error)
	GetMYSQLConnection(conns string) (*sql.DB, error)
	CloseConnection()
}

type mysqlService struct {
	conn       *sql.DB
	baseConfig map[string]*config.DBSQLConf
}

func NewService(dbConfig map[string]*config.DBSQLConf) MysqlService {
	return &mysqlService{
		conn:       db,
		baseConfig: dbConfig,
	}
}

func (s *mysqlService) GetMYSQLConnection(conns string) (*sql.DB, error) {
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
		for _, dbSession := range dbMap {
			dbSession.Close()
		}
	}
}

func mysqlConnect(connections string, dbConfig *config.DBSQLConf) (*sql.DB, error) {
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
	// max open connections
	dbMaxOpenConns, _ := strconv.Atoi(dbConfig.DbMaxOpenConns)
	MysqlDb.SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections
	dbMaxIdleConns, _ := strconv.Atoi(dbConfig.DbMaxIdleConns)
	MysqlDb.SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <=0 will forever
	dbMaxLifetimeConns, _ := strconv.Atoi(dbConfig.DbMaxLifetimeConns)
	MysqlDb.SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	if err = MysqlDb.Ping(); nil != err {
		color.Danger.Println(connections+"MySQL数据库连接失败。。。", err.Error())
		// os.Exit(200)
	} else {
		color.Info.Println(connections + "MySQL已连接 >>> ")
	}
	return MysqlDb, nil
}

func (s *mysqlService) InitConnection() (map[string]*sql.DB, error) {
	dbMaps := make(map[string]*sql.DB)
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
