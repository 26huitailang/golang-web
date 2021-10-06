package api

import (
	"fmt"
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/app/service/downloadsuite"
	"github.com/26huitailang/golang_web/config"
	"github.com/26huitailang/golang_web/library/response"
	"github.com/26huitailang/golang_web/middleware"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// API管理对象
var TaskThemeApi = new(taskThemeApi)
var TaskSuiteApi = new(taskSuiteApi)

type taskThemeApi struct{}
type taskSuiteApi struct{}

// @summary theme api
// @tags    theme
// @produce json
// @param   entity query model.SuitesQuery true "注册请求"
// @router  /apiV1/suites [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (t *taskThemeApi) Post(c echo.Context) (err error) {
	cc := c.(*middleware.CustomContext)
	log.Infof("user visited: %v", cc.Session)
	req := new(model.ReqDownloadTask)
	if err = c.Bind(req); err != nil {
		return
	}
	log.Println("url:", req.Url)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("%v", err)
			}
		}()
		t := downloadsuite.NewTheme(req.Url, config.Config.MediaPath)
		t.DownloadOneTheme()
		fmt.Printf("%v", t)
		downloadsuite.InitTheme(config.Config)
	}()
	return response.Json(c, response.OK, "task theme sent ...")
}

// @summary theme api
// @tags    theme
// @produce json
// @param   entity query model.SuitesQuery true "注册请求"
// @router  /apiV1/suites [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (t *taskSuiteApi) Post(c echo.Context) (err error) {
	cc := c.(*middleware.CustomContext)
	log.Infof("user visited: %v", cc.Session)
	req := new(model.ReqDownloadTask)
	if err = c.Bind(req); err != nil {
		return
	}
	log.Println("url:", req.Url)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("%v", err)
			}
		}()
		operator := downloadsuite.NewMeituriSuite(req.Url, config.Config.MediaPath, downloadsuite.MeituriParser{})
		suite := downloadsuite.NewSuite(operator)
		suite.Download()
		// 重新加载进去
		downloadsuite.InitTheme(config.Config)
	}()
	return response.Json(c, response.OK, "task downloadsuite sent ...")
}
