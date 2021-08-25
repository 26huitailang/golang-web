package api

import (
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/library/response"
	"github.com/labstack/echo"
)

// @summary theme list api
// @tags    theme service
// @produce json
// @param   entity  body model.UserApiSignUpReq true "注册请求"
// @router  /suites [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func Login(c echo.Context) (err error) {
	req := new(model.ApiLoginReq)
	if err = c.Bind(req); err != nil {
		return
	}

	return response.Json(c, 1, "ok")
}
