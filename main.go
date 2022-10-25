package main

import (
	"server.pluto.com/log"
	"server.pluto.com/module"
	"server.pluto.com/mysql"
)

func main() {
	mysql.InitMysql()
	log.SetLevel(log.InfoLevel)

	log.InitLogToFile()

	err := module.RegisterRouter(module.NewRouter()).Run(":7777")

	if err != nil {
		log.Error(err)
	}
}
