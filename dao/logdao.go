package dao

import (
	"mail_log/entity"
	"mail_log/util"
	"time"
)

var LogItem []string = []string{"INFO", "ERROR", "USER", "SELF"}

type LogDao struct {
	db *util.Database
}

func (dao *LogDao) Init(db *util.Database) {
	dao.db = db
}

func (dao LogDao) Write(logEntity *entity.LogEntity) {
	table := logEntity.Level + time.Now().Format("200601")
	dao.db.DB.Table(table).AutoMigrate(&entity.LogEntity{})
	dao.db.DB.AutoMigrate(&entity.Item{})
	var count int64
	dao.db.DB.Model(&entity.Item{}).Where("Name = ?", table).Count(&count)
	if count == 0 {
		dao.db.DB.Create(&entity.Item{
			Name: table,
		})
	}
	dao.db.DB.Table(table).Create(logEntity)
}
