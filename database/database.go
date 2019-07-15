package database

import (
	"log"

	"golang_web/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init() {
	// DB 小心:= 覆盖了声明的全局变量
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Panicf("DB connect error: %s", err)
	}
	DB.LogMode(true)

	// 迁移
	DB.SingularTable(true) // 单数表名
	DB.AutoMigrate(&models.Theme{}, &models.Suite{}, &models.Image{})
	// sqlite 对alter table的支持有限，不支持rename column和remove column
	// err = DB.Model(&Image{}).DropColumn("IsRead").Erro
}
