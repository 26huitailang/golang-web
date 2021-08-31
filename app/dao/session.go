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
	ret := d.DB().Where(&model.Session{Token: token}).First(session)
	if ret.RowsAffected == 0 {
		return nil
	}
	return session
}

func (d *sessionDao) DeleteOne(token string) (rows int) {
	session := d.GetOne(token)
	ret := d.DB().Delete(session)
	rows = int(ret.RowsAffected)
	return rows
}
