package main

import (
	"mail_log/dao"
	"mail_log/util"
)

func main() {
	db := &util.Database{}
	db.Init()
	logdao := &dao.LogDao{}
	logdao.Init(db)
}
