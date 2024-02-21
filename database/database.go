package database

import (
	"log"
	"path"
	"strings"

	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type DatabaseStore struct{}

func init() {
	//db = connectDB(db, config.Config.DB)
}

func NewDatabaseStore() *DatabaseStore {
	return &DatabaseStore{}
}

func (s *DatabaseStore) DB() *gorm.DB {
	return db
}

func tables() []interface{} {
	return []interface{}{
		&model.Session{},
		&model.Theme{},
		&model.Suite{},
		&model.Image{},
		&model.User{},
	}
}

func connectDB(database *gorm.DB, dbFile string) *gorm.DB {
	// DB 小心:= 覆盖了声明的全局变量
	var err error
	database, err = gorm.Open("sqlite3", path.Join(config.Config.DataPath, dbFile))
	if err != nil {
		log.Panicf("DB connect error: %s", err)
	}
	database.LogMode(false)

	// 迁移
	database.SingularTable(true) // 单数表名
	database.CreateTable(tables()...)
	database.AutoMigrate(tables()...)
	// sqlite 对alter table的支持有限，不支持rename column和remove column
	return database
}

func DropTables(db *gorm.DB) {
	db.DropTable(tables()...)
}

func TestDB() *gorm.DB {
	dbFile := strings.Join([]string{config.Config.DB, "test"}, "")
	db = connectDB(db, dbFile)
	return db
}
