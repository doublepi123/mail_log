package util

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) Init() {
	var err error
	db.DB, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

}
