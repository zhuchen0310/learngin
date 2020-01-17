package database

import (
	// database driver

	"log"

	"github.com/zhuchen/learngin/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// NewMySQL new db and retry connection when has error.
func NewMySQL() *gorm.DB {
	if db != nil {
		log.Panicln("db 单例")
		return db
	}
	db, err := gorm.Open("mysql", config.MysqlDsn)
	if err != nil {
		log.Panicln("db error: ", err.Error())
	}
	return db
}
