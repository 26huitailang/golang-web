package router

import (
	"fmt"

	"github.com/26huitailang/golang_web/app/api"
	"github.com/26huitailang/golang_web/app/service"
	"github.com/26huitailang/golang_web/config"
	_ "github.com/26huitailang/golang_web/docs"
	"github.com/26huitailang/golang_web/library/response"
	"github.com/26huitailang/golang_web/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SessionCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := c.Cookie("token")
		if err != nil {
			return response.Json(c, response.AuthCookieInvalid, fmt.Sprintf("get cookie error: %s", err.Error()))
		}
		sessionVal := service.UserService.GetSession(token.Value)
		if sessionVal == nil {
			return response.Json(c, response.AuthCookieExpired, "null cookie")
		}
		cc := c.(*middleware.CustomContext)
		session := service.UserService.GetSession(token.Value)
		cc.Session = session
		return next(cc)
	}
}

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

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/login", api.Login)
	e.POST("/logout", api.Logout, SessionCheckMiddleware)
	g := e.Group("/apiV1")
	g.Use(SessionCheckMiddleware)
	g.GET("/suites", api.SuiteRest.Get)
	//e.GET("/suites/:suite_id", views.SuiteHandle)
	//e.GET("/suites/:id/doread", views.SuiteReadHandle)
	//e.GET("/suites/:id/dolike", views.SuiteLikeHandle)

	//devopsGroup := e.Group("/devops")
	//devopsGroup.POST("/initdb", views.InitDBHandle)
	//devopsGroup.GET("", views.DevopsHandle)
	e.Static("/image/*filepath", config.Config.MediaPath)
}
