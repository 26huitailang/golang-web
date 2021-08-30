package server

import (
	"github.com/26huitailang/golang_web/app/api"
	"github.com/26huitailang/golang_web/config"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	// profiling
	// mux = httpprof.WrapRouter(mux)
	// e.HTTPErrorHandler = customHTTPErrorHandler
	//e.GET("/", views.IndexHandle)
	//e.GET("/ws_view/ws", ws_hello)
	//e.GET("/ws_view", ws_view)
	//e.GET("/hello/:name", func(c echo.Context) error {
	//	name := c.Param("name")
	//	resp := fmt.Sprintf("Hello, %s!", name)
	//	return c.String(http.StatusOK, resp)
	//})

	//e.POST("/task/suite", views.TaskSuiteHandle)
	//e.POST("/task/theme", views.TaskThemeHandle)

	//e.GET("/themes", handler.ThemesHandle)
	//e.GET("/themes/:id", views.ThemeHandle)
	e.POST("/login", api.Login)
	e.GET("/suites", api.SuiteRest.Get)
	//e.GET("/suites/:suite_id", views.SuiteHandle)
	//e.GET("/suites/:id/doread", views.SuiteReadHandle)
	//e.GET("/suites/:id/dolike", views.SuiteLikeHandle)

	//devopsGroup := e.Group("/devops")
	//devopsGroup.POST("/initdb", views.InitDBHandle)
	//devopsGroup.GET("", views.DevopsHandle)
	e.Static("/image/*filepath", config.Config.MediaPath)
}
