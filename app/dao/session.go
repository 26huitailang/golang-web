package dao

import (
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"
)

var Session = &sessionDao{
	&database.DatabaseStore{},
}

type sessionDao struct {
	*database.DatabaseStore
}

func (d *sessionDao) CreateOne(session *model.Session) (*model.Session, error) {
	result := d.DB().Create(session)
	return session, result.Error
}

func (d *sessionDao) GetOne(token string) (session *model.Session) {
	session = &model.Session{}
	d.DB().Where(&model.Session{Token: token}).First(session)
	return session
}
