package api

import (
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"
	"github.com/26huitailang/golang_web/library/response"
	"github.com/labstack/echo"
)

// API管理对象
var SuiteRest = new(suiteRestApi)

type suiteRestApi struct{}

// @summary theme list api
// @tags    theme service
// @produce json
// @param   entity  body model.UserApiSignUpReq true "注册请求"
// @router  /suites [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *suiteRestApi) Get(c echo.Context) (err error) {
	query := new(model.SuitesQuery)
	if err = c.Bind(query); err != nil {
		return
	}
	var suites []model.Suite

	DB := database.NewDatabaseStore().DB()
	DB.Where("is_like = ?", query.IsLike).Find(&suites)
	return response.Json(c, 1, "ok", suites)
}
