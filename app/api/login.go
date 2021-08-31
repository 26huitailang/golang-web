package api

import (
	"encoding/json"
	"fmt"
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
		return response.Json(c, response.ReqParamInvalid, "invalid req params!")
	}
	if err = c.Validate(req); err != nil {
		return response.Json(c, response.ReqParamInvalid, fmt.Sprintf("validated faild: %v", err.Error()))
	}
	ok, user := service.UserService.Authenticate(req.Username, req.Password)
	if !ok {
		return response.Json(c, response.AuthFailed, "authenticate failed!")
	}

	valBytes, err := json.Marshal(&model.SessionValue{ID: user.ID, Username: user.Username, Nickname: user.Nickname})
	if err != nil {
		return response.Json(c, response.AuthCreateSessionFailed, fmt.Sprintf("create session failed: %s", err.Error()))
	}

	token := service.UserService.CreateSession(string(valBytes))
	c.SetCookie(&http.Cookie{Name: "token", Value: token, HttpOnly: true})
	return response.Json(c, response.OK, "ok", map[string]string{"token": token})
}
