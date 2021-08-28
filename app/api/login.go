package api

import (
	"net/http"

	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/app/service"
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
		return response.Json(c, -1, "invalid req params")
	}
	if ok := service.UserService.Authenticate(req.Username, req.Password); !ok {
		return response.Json(c, -1, "authenticate failed!")
	}
	c.SetCookie(&http.Cookie{Name: "Token", Value: "hello"})
	return response.Json(c, response.OK, "ok")
}
