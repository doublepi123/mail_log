package dao

import (
	"mail_log/entity"
	"mail_log/util"
)

var LogItem []string = []string{"INFO", "ERROR", "USER", "SELF"}

type LogDao struct {
	db *util.Database
}

func (dao *LogDao) Init(db *util.Database) {
	dao.db = db
	for i := range LogItem {
		dao.db.DB.Table(LogItem[i]).AutoMigrate(&entity.LogEntity{})
	}
	var str string
	for i := range LogItem {
		str += LogItem[i] + ","
	}
	dao.db.Redis.Set("items", str, 0)
}

func (dao LogDao) Write(logEntity *entity.LogEntity) {
	dao.db.DB.Table(logEntity.Level).Create(logEntity)
}
