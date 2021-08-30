package dao

import (
	"fmt"
	"testing"
	"time"

	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"
	"github.com/26huitailang/golang_web/utils"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SessionTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *SessionTestSuite) SetupTest() {
	fmt.Println("Setup")
	dbStore := &database.DatabaseStore{}
	database.TestDB()
	suite.db = dbStore.DB()
}

func (suite *SessionTestSuite) TearDownTest() {
	fmt.Println("TearDown")
	database.DropTables(suite.db)
}

func (suite *SessionTestSuite) TestSession_GetOne() {
	suite.T().Run("get one ok", func(t *testing.T) {

		item := &model.Session{Token: utils.UUID(), Value: "hello", ExpiredAt: time.Now()}
		_, _ = Session.CreateOne(item)
		got := Session.GetOne(item.Token)
		assert.Equal(t, item.Token, got.Token)
		assert.Equal(t, item.Value, got.Value)
		assert.Equal(t, item.ExpiredAt, got.ExpiredAt) // TODO: 时间存入取出不一致
	})
}

func TestSessionTestSuite(t *testing.T) {
	suite.Run(t, new(SessionTestSuite))
}
