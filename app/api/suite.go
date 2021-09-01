package api

import (
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"
	"github.com/26huitailang/golang_web/library/response"
	"github.com/26huitailang/golang_web/middleware"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// API管理对象
var SuiteRest = new(suiteRestApi)

type suiteRestApi struct{}

// @summary suite api
// @tags    suite
// @produce json
// @param   entity query model.SuitesQuery true "注册请求"
// @router  /apiV1/suites [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *suiteRestApi) Get(c echo.Context) (err error) {
	cc := c.(*middleware.CustomContext)
	log.Infof("user visited: %v", cc.Session)
	query := new(model.SuitesQuery)
	if err = c.Bind(query); err != nil {
		return
	}
	var suites []model.Suite

	DB := database.NewDatabaseStore().DB()
	DB.Where("is_like = ?", query.IsLike).Find(&suites)
	return response.Json(c, 1, "ok", suites)
}
