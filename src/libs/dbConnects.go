package libs


import (
	"configs"
	_ "github.com/go-sql-driver/mysql"
)

type mysqlConfig struct {
	name 		string
	port 		int
	host 		string
	user 		string
	password	string
	charset     string

}

var (
	mysqlHost = configs.LoadConfig("TeCentYun")["host"]
	mysqlPort = configs.LoadConfig("TeCentYun")["port"]
	mysqlUser = configs.LoadConfig("TeCentYun")["user"]
	mysqlPassword = configs.LoadConfig("TeCentYun")["password"]
	mysqlCharset = configs.LoadConfig("TeCentYun")["charset"]
)

func connectMysql() {
}








