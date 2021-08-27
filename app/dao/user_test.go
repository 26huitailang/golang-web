package dao

import (
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *UserTestSuite) SetupTest() {
	suite.db = database.TestDB()
	User.db = suite.db
}

// TODO: 完成这个测试suite

func (suite *UserTestSuite) TestUserDao_GetOne() {
	suite.Run("get one ok", func(t *testing.T) {
		user := &model.User{Username: "test", Password: "123123"}
		println(User.db)
		_, createdUser := User.CreateOne(user)
		got := User.GetOne(createdUser.ID)
		assert.Equal(t, user.Username, got.Username)
	})
	suite.Run("get one ok", func(t *testing.T) {
		user := model.User{Username: "test", Password: "123123"}
		_, createdUser := User.CreateOne(&user)
		got := User.GetOne(createdUser.ID)
		assert.Equal(t, user.Username, got.Username)
	})
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
