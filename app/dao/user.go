package dao

import (
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"
)

var User = &userDao{
	DatabaseStore: &database.DatabaseStore{},
}

type userDao struct {
	*database.DatabaseStore
}

func (d *userDao) CreateOne(User *model.User) (*model.User, error) {
	result := d.DB().Create(User)
	return User, result.Error
}

func (d *userDao) GetOne(Id uint) (user *model.User) {
	user = &model.User{}
	d.DB().First(user, Id)
	return user
}

func (d *userDao) GetOneByUsername(username string) *model.User {
	user := &model.User{}
	d.DB().Where(&model.User{Username: username}).First(user)
	return user
}
