package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/26huitailang/golang_web/app/service"

	"github.com/26huitailang/golang_web/app/api"
	"github.com/26huitailang/golang_web/app/dao"
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"
	"github.com/26huitailang/golang_web/library/mycrypto"
	"github.com/26huitailang/golang_web/library/response"
	"github.com/26huitailang/golang_web/server"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
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
	suite.e = server.NewServer()
}

func (suite *LoginTestSuite) TearDownTest() {
	fmt.Println("TearDown")
	database.DropTables(suite.db)
}

func (suite *LoginTestSuite) TestLogin() {
	suite.T().Run("test login ok", func(t *testing.T) {
		userJSON := `{"username": "test", "password": "123123"}`
		dao.User.CreateOne(&model.User{Username: "test", Nickname: "nickname", Password: mycrypto.Password("123123").Encrypt(nil)})
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
			data := resp.Data.(map[string]interface{})
			sessionVal := service.UserService.GetSession(data["token"].(string))
			assert.Equal(t, "test", sessionVal.Username)
			assert.Equal(t, "nickname", sessionVal.Nickname)
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
			assert.Equal(t, response.AuthFailed, resp.Code)
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
			assert.Equal(t, response.ReqParamInvalid, resp.Code)
			assert.GreaterOrEqual(t, "validated failed:", resp.Message)
		}
	})
}

func (suite *LoginTestSuite) TestLogout() {
	suite.T().Run("logout ok", func(t *testing.T) {
		userJSON := `{"username": "test", "password": "123123"}`
		dao.User.CreateOne(&model.User{Username: "test", Nickname: "nickname", Password: mycrypto.Password("123123").Encrypt(nil)})
		token := service.UserService.CreateSession(userJSON)
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(userJSON))
		req.AddCookie(&http.Cookie{Name: "token", Value: token})
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := suite.e.NewContext(req, rec)
		c.SetPath("/logout")
		c.SetRequest(req)

		// Assertions
		if assert.NoError(t, api.Logout(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var resp response.JsonResponse
			_ = json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.Equal(t, response.OK, resp.Code)
			assert.Equal(t, "logout succeed", resp.Message)
			sessionVal := service.UserService.GetSession(token)
			assert.Equal(t, (*model.SessionValue)(nil), sessionVal)
		}
	})
}
func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
