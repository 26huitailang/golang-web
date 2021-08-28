package internal

import (
	"github.com/26huitailang/golang_web/app/dao"
	"github.com/26huitailang/golang_web/app/model"
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
