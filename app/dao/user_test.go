package dao

import (
	"fmt"
	"testing"

	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *UserTestSuite) SetupTest() {
	fmt.Println("Setup")
	suite.db = database.TestDB()
}

func (suite *UserTestSuite) TearDownTest() {
	fmt.Println("TearDown")
	database.DropTables(suite.db)
}

func (suite *UserTestSuite) TestUserDao_GetOne() {
	suite.T().Run("get one ok", func(t *testing.T) {
		user := &model.User{Username: "test", Password: "123123"}
		createdUser, _ := User.CreateOne(user)
		got := User.GetOne(createdUser.ID)
		assert.Equal(t, user.Username, got.Username)
	})
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
