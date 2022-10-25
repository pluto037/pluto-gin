package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"server.pluto.com/log"
)

var (
	username  = "root"
	password  = "root"
	ipAddress = "127.0.0.1"
	port      = 3306
	dbNmae    = "mc"
	charset   = "utf-8"
)

var Engine *xorm.Engine

var err error

func InitMysql() {
	url := "root:root@tcp(127.0.0.1)/mc?charset=utf8"
	// url := fmt.Sprint("%s:%s@tcp(%s)/%s?charset=%s", username, password, ipAddress, dbNmae, charset)
	eng, err := xorm.NewEngine("mysql", url)
	if err != nil {
		log.Error("数据库连接失败", err)
	}
	Engine = eng

}
