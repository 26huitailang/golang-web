package internal

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/26huitailang/golang_web/app/dao"
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/config"
	"github.com/26huitailang/golang_web/library/mycrypto"
	"github.com/26huitailang/golang_web/library/utils"
)

var UserService = &userService{}

type userService struct{}

func (s *userService) CreateUser(user *model.User) (*model.User, error) {
	return dao.User.CreateOne(user)
}

func (s *userService) GetUser(username string) *model.User {
	return dao.User.GetOneByUsername(username)
}

func (s *userService) Authenticate(username, password string) (bool, *model.User) {
	user := dao.User.GetOneByUsername(username)
	if user == nil {
		return false, nil
	}
	pwd := mycrypto.Password(password)
	return pwd.Check(user.Password), user
}

func (s *userService) CreateSession(value string) string {
	ExpiredAt := time.Now().Add(time.Second * time.Duration(config.Config.SessionExpiredTime))
	session := &model.Session{
		Token:     utils.UUID(),
		Value:     value,
		ExpiredAt: ExpiredAt,
	}
	session, err := dao.Session.CreateOne(session)
	if err != nil {
		return ""
	}
	return session.Token
}

func (s *userService) GetSession(token string) *model.SessionValue {
	session := dao.Session.GetOne(token)
	if session == nil {
		return nil
	}
	if session.ExpiredAt.Before(time.Now()) {
		log.Warningf("session unmarshal failed: %s %v", session.Value, session.ExpiredAt)
		return nil
	}
	var sessionVal model.SessionValue
	if err := json.Unmarshal([]byte(session.Value), &sessionVal); err != nil {
		log.Errorf("session unmarshal failed: %s %v", err.Error(), session.Value)
		return nil
	}
	return &sessionVal
}

func (s *userService) Logout(token string) {
	_ = dao.Session.DeleteOne(token)
}
