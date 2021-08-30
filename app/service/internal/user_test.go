package internal

import (
	"fmt"
	"testing"

	"github.com/26huitailang/golang_web/app/dao"
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"
	"github.com/26huitailang/golang_web/utils/mycrypto"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *AuthTestSuite) SetupTest() {
	fmt.Println("Setup")
	suite.db = database.TestDB()
	dao.User.DB = database.TestDB()
}
func (suite *AuthTestSuite) TearDownTest() {
	fmt.Println("TearDown")
	database.DropTables(suite.db)
}

func tableTestSetup() func() {
	dao.User.DB = database.TestDB()
	return func() {
		database.DropTables(database.TestDB())
	}
}

func (suite *AuthTestSuite) TestuserService_Authenticate() {
	suite.T().Run("user not existed", func(t *testing.T) {
		passwd := UserService.Authenticate("test", "hello")
		suite.Assert().Equal(false, passwd)
	})
	suite.T().Run("user auth ok", func(t *testing.T) {
		type args struct {
			Username string
			Password string
			ReqPwd   string
		}
		testCases := []struct {
			desc string
			args args
			want bool
		}{
			{desc: "wrong password not passed", args: args{Username: "test", Password: "123123", ReqPwd: "321321"}, want: false},
			{desc: "pwd: 123123", args: args{Username: "test", Password: "123123", ReqPwd: "123123"}, want: true},
			{desc: "pwd: hello!@#", args: args{Username: "test", Password: "hello!@#", ReqPwd: "hello!@#"}, want: true},
			{desc: "pwd: 中文测试", args: args{Username: "test", Password: "中文测试", ReqPwd: "中文测试"}, want: true},
		}
		for _, tC := range testCases {
			t.Run(tC.desc, func(t *testing.T) {
				tbTestTD := tableTestSetup()
				defer tbTestTD()

				user := &model.User{
					Username: tC.args.Username,
					Password: mycrypto.Password(tC.args.Password).Encrypt(nil),
				}
				UserService.CreateUser(user)
				passwd := UserService.Authenticate(tC.args.Username, tC.args.ReqPwd)
				suite.Assert().Equal(tC.want, passwd, tC.desc)
			})
		}
	})
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
