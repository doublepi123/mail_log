package entity

import "gorm.io/gorm"

type LogEntity struct {
	gorm.Model
	WorkID  string
	Name    string
	Level   string
	Message string
}

type PictureEntity struct {
	gorm.Model
	OrgName  string
	Name     string
	FileName string
	Data     []byte
	Pid      string
}

type Item struct {
	gorm.Model
	Name string `gorm:"unique"`
}

func LogTable(log LogEntity) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table(log.Level)
	}
}
