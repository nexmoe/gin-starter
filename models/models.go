package models

import (
	"log"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"store-server/pkg/setting"
)

var db *gorm.DB

func init() {
	var (
		err error
	)

	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.Cfg.GetString("mysql.username"),
		setting.Cfg.GetString("mysql.password"),
		setting.Cfg.GetString("mysql.host"),
		setting.Cfg.GetString("mysql.dbname")))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	    return setting.Cfg.GetString("mysql.table_prefix") + defaultTableName;
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}