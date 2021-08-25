package auth

var UserService = userService{}

type userService struct{}

//func (s *userService) CreateUser(username, password string) (err error, user model.User) {
//	user.Username = username
//	binPwd := pbkdf2.Key([]byte(password), []byte("123123"), 1, 32, sha256.New)
//
//	user.Password =	string(binPwd)
//}

func (s userService) Authenticate(username, password string) (err error) {
	//	todo: 完成测试和设计
	return nil
}
