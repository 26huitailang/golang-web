package database

import (
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"path"
	"strings"
)

var db *gorm.DB
var testDb *gorm.DB

func init() {
	connectDB(db, config.Config.DB)
}

func DB() *gorm.DB {
	return db
}

func connectDB(db *gorm.DB, dbFile string) {
	// DB 小心:= 覆盖了声明的全局变量
	var err error
	db, err = gorm.Open("sqlite3", path.Join(config.Config.DataPath, dbFile))
	if err != nil {
		log.Panicf("DB connect error: %s", err)
	}
	db.LogMode(true)

	// 迁移
	db.SingularTable(true) // 单数表名
	db.CreateTable(&model.User{})
	db.AutoMigrate(&model.Theme{}, &model.Suite{}, &model.Image{}, &model.User{})
	// sqlite 对alter table的支持有限，不支持rename column和remove column
	// err = DB.Model(&Image{}).DropColumn("IsRead").Erro
}

func TestDB() *gorm.DB {
	dbFile := strings.Join([]string{config.Config.DB, "test"}, "")
	connectDB(testDb, dbFile)
	return testDb
}
