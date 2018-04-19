package models

import (
	"configs"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"github.com/prometheus/common/log"
)

type BaseDb struct {
	dbName 		string
	tbName		string
	db 			*sql.DB
}

var (
	mysqlHost = configs.LoadConfig("TeCentYun")["host"]
	mysqlPort = configs.LoadConfig("TeCentYun")["port"]
	mysqlUser = configs.LoadConfig("TeCentYun")["user"]
	mysqlPassword = configs.LoadConfig("TeCentYun")["password"]
	mysqlCharset = configs.LoadConfig("TeCentYun")["charset"]
)


func NewBaseDb(dbName string, tbName string) *BaseDb {
	bd := &BaseDb{}
	bd.dbName = dbName
	bd.tbName = tbName
	bd.db = connectMysql()
	log.Info("init BaseDb %v - %v", bd.dbName, bd.tbName)
	return bd
}

//连接数据库
func connectMysql() *sql.DB {
	//db, err := sql.Open("mysql", "mysqlUser:mysqlPassword@tcp(mysqlHost:mysqlPort)/Test?mysqlCharset")
	db, err := sql.Open("mysql",fmt.Sprintf("%v:%v@tcp(%v:%v)/test?%v",mysqlUser,mysqlPassword,mysqlHost,mysqlPort,mysqlCharset))
	if err != nil {
		log.Error("sql Open is falied err = %v", err)
		return nil
	}
	return db
}








