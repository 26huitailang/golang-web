package database

import (
	"golang_web/config"
	"log"
	"path"

	"golang_web/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	// DB 小心:= 覆盖了声明的全局变量
	var err error
	db, err = gorm.Open("sqlite3", path.Join(config.Config.DataPath, "test.db"))
	if err != nil {
		log.Panicf("DB connect error: %s", err)
	}
	db.LogMode(true)

	// 迁移
	db.SingularTable(true) // 单数表名
	db.AutoMigrate(&models.Theme{}, &models.Suite{}, &models.Image{})
	// sqlite 对alter table的支持有限，不支持rename column和remove column
	// err = DB.Model(&Image{}).DropColumn("IsRead").Erro
}

func DB() *gorm.DB {
	return db
}
