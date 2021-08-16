package driver

//import (
//	"database/sql"
//	"math/rand"
//)
//import "order/config"

//type DB struct {
//	masterDB *sql.DB
//	slaveDB  []*sql.DB
//	Config   *config
//}
//
//func GetDSN(conn string) string {
//	cfg := config.GetMySQLConfig()
//	dsn := cfg["user"] + ":" + cfg["password"] + "@tcp(" + cfg["host"] + ":" + cfg["port"] + ")/" + cfg["db"] + "?charset=" + cfg["charset"]
//	return dsn
//}
//
//func New(c *Config) (db *DB, err error) {
//	db = new(DB)
//	db.Config = c
//	db.masterDB, err = sql.Open("mysql", c.Master.DSN)
//	if err != nil {
//		err = errorsWrap(err, "init master db error")
//		return
//	}
//
//	db.masterDB.SetMaxOpenConns(c.Master.MaxOpen)
//	db.masterDB.SetMaxIdleConns(c.Master.MaxIdle)
//	if err = db.masterDB.Ping(); err != nil {
//		err = errorsWrap(err, "master db ping error")
//		return
//	}
//
//	for i := 0; i < len(c.Slave); i++ {
//		var mysqlDB *sql.DB
//		mysqlDB, err = sql.Open("mysql", c.Slave[i].DSN)
//		if err != nil {
//			err = errorsWrap(err, "init slave db error")
//			return
//		}
//
//		mysqlDB.SetMaxOpenConns(c.Slave[i].MaxOpen)
//		mysqlDB.SetMaxIdleConns(c.Slave[i].MaxIdle)
//		if err = mysqlDB.Ping(); err != nil {
//			err = errorsWrap(err, "slave db ping error")
//			return
//		}
//
//		db.slaveDB = append(db.slaveDB, mysqlDB)
//	}
//	return
//}
//
//func (db *DB) MasterDB() *sql.DB {
//	return db.masterDB
//}
//
//func (db *DB) SlaveDB() *sql.DB {
//	if len(db.slaveDB) == 0 {
//		return db.masterDB
//	}
//	n := rand.Intn(len(db.slaveDB))
//	return db.slaveDB[n]
//}
//
//// MasterDBClose 释放主库的资源
//func (db *DB) MasterDBClose() error {
//	if db.masterDB != nil {
//		return db.masterDB.Close()
//	}
//	return nil
//}
//
//// SlaveDBClose 释放从库的资源
//func (db *DB) SlaveDBClose() (err error) {
//	for i := 0; i < len(db.slaveDB); i++ {
//		err = db.slaveDB[i].Close()
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
