package server

import (
	"fmt"
	"golang_web/config"
	"golang_web/constants"
	"golang_web/database"
	"golang_web/views"

	// "log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/net/websocket"
)

// 初始化文件结构
func Init() {
	log.SetLevel(log.DebugLevel)
	// 模板预加载
	config.ReloadTemplates()
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	log.Errorf("status code: %d", code)
	errorPage := fmt.Sprintf("templates/error/%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
}

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) SetConfig() {
	log.Infoln("setting config...")
	c.Set("config", config.Config)
	log.Infoln("finish set config!")
}

func ws_hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func ws_view(c echo.Context) error {
	return c.Render(200, "layout:websocket", "")
}

func NewServer() *echo.Echo {
	Init()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			cc.SetConfig()
			println("CCCCCCC:", cc.Get("config").(*config.Configuration).Port)
			return h(cc)
		}
	})

	if config.Config.DeployLevel <= constants.Development {
		e.Use(middleware.Logger())
	} else {
		logfile, _ := os.OpenFile("main.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: logfile,
		}))
	}
	// e.Use(middleware.CSRF())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	// e.Use(middleware.JWT([]byte("secret")))

	var EchoTemplate = &config.Template{}
	e.Renderer = EchoTemplate

	DB := database.DB()
	store := &views.DatabaseStore{DB: DB}
	handler := &views.Handler{Store: store}
	// profiling
	// mux = httpprof.WrapRouter(mux)
	// e.HTTPErrorHandler = customHTTPErrorHandler
	e.GET("/", views.IndexHandle)
	e.GET("/ws_view/ws", ws_hello)
	e.GET("/ws_view", ws_view)
	e.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")
		resp := fmt.Sprintf("Hello, %s!", name)
		return c.String(http.StatusOK, resp)
	})

	e.POST("/task/suite", views.TaskSuiteHandle)
	e.POST("/task/theme", views.TaskThemeHandle)

	e.GET("/themes", handler.ThemesHandle)
	e.GET("/themes/:id", views.ThemeHandle)
	e.GET("/suites", views.SuitesHandle)
	e.GET("/suites/:suite_id", views.SuiteHandle)
	e.GET("/suites/:id/doread", views.SuiteReadHandle)
	e.GET("/suites/:id/dolike", views.SuiteLikeHandle)

	devopsGroup := e.Group("/devops")
	devopsGroup.POST("/initdb", views.InitDBHandle)
	devopsGroup.GET("", views.DevopsHandle)
	e.Static("/image/*filepath", config.Config.MediaPath)

	addr := fmt.Sprintf("%s%s", config.Config.IP, config.Config.Port)
	fmt.Printf("serve: http://%s\n", addr)

	return e
}
