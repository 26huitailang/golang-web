package internal

import (
	"time"

	"github.com/26huitailang/golang_web/app/dao"
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/config"
	"github.com/26huitailang/golang_web/utils"
	"github.com/26huitailang/golang_web/utils/mycrypto"
)

var UserService = userService{}

type userService struct{}

func (s *userService) CreateUser(user *model.User) (*model.User, error) {
	return dao.User.CreateOne(user)
}

func (s *userService) Authenticate(username, password string) bool {
	user := dao.User.GetOneByUsername(username)
	if user == nil {
		return false
	}
	pwd := mycrypto.Password(password)
	return pwd.Check(user.Password)
}

// TODO do tests
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
