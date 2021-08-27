package dao

import (
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"
	"github.com/jinzhu/gorm"
)

var User = userDao{
	db: database.DB(),
}

type userDao struct {
	db *gorm.DB
}

func (d userDao) CreateOne(User *model.User) (error, *model.User) {
	result := d.db.Create(User)
	return result.Error, User
}
func (d userDao) GetOne(Id uint) (user model.User) {
	d.db.First(user, Id)
	return user
}
