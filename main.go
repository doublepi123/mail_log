package main

import (
	"mail_log/dao"
	"mail_log/service"
	"mail_log/util"
)

func main() {
	db := &util.Database{}
	db.Init()
	logdao := &dao.LogDao{}
	logdao.Init(db)
	logsaver := service.LogSaverServer{}
	logsaver.Init(logdao, db.Redis)
	logsaver.Run()
}
