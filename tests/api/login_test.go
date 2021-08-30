package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/26huitailang/golang_web/app/api"
	"github.com/26huitailang/golang_web/app/dao"
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"
	"github.com/26huitailang/golang_web/library/response"
	"github.com/26huitailang/golang_web/server"
	"github.com/26huitailang/golang_web/utils/mycrypto"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoginTestSuite struct {
	suite.Suite
	db *gorm.DB
	e  *echo.Echo
}

func (suite *LoginTestSuite) SetupTest() {
	fmt.Println("Setup")
	suite.db = database.TestDB()
	dao.User.DB = database.TestDB()
	suite.e = server.NewServer()
}

func (suite *LoginTestSuite) TearDownTest() {
	fmt.Println("TearDown")
	database.DropTables(suite.db)
}

func (suite *LoginTestSuite) TestLogin() {
	suite.T().Run("test login ok", func(t *testing.T) {
		userJSON := `{"username": "test", "password": "123123"}`
		dao.User.CreateOne(&model.User{Username: "test", Password: mycrypto.Password("123123").Encrypt(nil)})
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := suite.e.NewContext(req, rec)
		c.SetPath("/login")
		c.SetRequest(req)

		// Assertions
		if assert.NoError(t, api.Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var resp response.JsonResponse
			_ = json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.Equal(t, response.OK, resp.Code)
			assert.Equal(t, "ok", resp.Message)
		}
	})
	suite.T().Run("test auth failed", func(t *testing.T) {
		userJSON := `{"username": "test", "password": "321321"}`
		dao.User.CreateOne(&model.User{Username: "test", Password: mycrypto.Password("123123").Encrypt(nil)})
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := suite.e.NewContext(req, rec)
		c.SetPath("/login")
		c.SetRequest(req)

		// Assertions
		if assert.NoError(t, api.Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var resp response.JsonResponse
			_ = json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.Equal(t, response.AUTH_FAILED, resp.Code)
			assert.Equal(t, "authenticate failed!", resp.Message)
		}
	})
	suite.T().Run("invalid params", func(t *testing.T) {
		userJSON := `{"kkkkkkkkkk": "test_error", "kllklkl": "123123"}`
		dao.User.CreateOne(&model.User{Username: "test_error", Password: mycrypto.Password("123123").Encrypt(nil)})
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := suite.e.NewContext(req, rec)
		c.SetPath("/login")
		c.SetRequest(req)

		// Assertions
		if assert.NoError(t, api.Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var resp response.JsonResponse
			_ = json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.Equal(t, response.REQ_PARAM_INVALID, resp.Code)
			assert.GreaterOrEqual(t, "validated failed:", resp.Message)
		}
	})
}

func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
