package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// ORM gorm引擎的实例
var ORM *gorm.DB

func InitDB(){
	//var err error
	//连接数据库，为方便教学，数据库使用sqlite
	db, err := gorm.Open("sqlite3", "./database/test.db")
	if err != nil {
		panic("连接数据库失败")
	}

	ORM = db
}