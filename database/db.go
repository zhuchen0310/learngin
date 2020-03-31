package database

import (
	// database driver

	"log"

	"github.com/zhuchen/learngin/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// init 包初始化函数类似Python init.py
func init() {
	var err error
	db, err = gorm.Open("mysql", config.MysqlDsn)
	if err != nil {
		log.Panicln("db error: ", err.Error())
	}
	db.DB().SetMaxIdleConns(10) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	db.DB().SetMaxOpenConns(50) //设置数据库连接池最大连接数
}

// NewMySQL 获取新的 mysql 连接
func NewMySQL() *gorm.DB {
	return db
}
