package main

import (
	"fmt"

	"golang_web/database"

	// "log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/net/websocket"
)

var DB = database.DB

// 初始化文件结构
func init() {
	// var err error
	// var config = config.Config
	// log
	// logfile, err := os.OpenFile("main.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	// 	log.SetOutput(logfile)
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	// initConfiguration()

	// 模板预加载
	ReloadTemplates()
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
	c.Set("config", Config)
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

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			cc.SetConfig()
			println("CCCCCCC:", cc.Get("config").(*Configuration).Port)
			return h(cc)
		}
	})

	if Config.DeployLevel >= Development {
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

	var EchoTemplate = &Template{}
	e.Renderer = EchoTemplate

	// profiling
	// mux = httpprof.WrapRouter(mux)
	// e.HTTPErrorHandler = customHTTPErrorHandler
	e.GET("/", IndexHandle)
	e.GET("/ws_view/ws", ws_hello)
	e.GET("/ws_view", ws_view)
	e.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")
		resp := fmt.Sprintf("Hello, %s!", name)
		return c.String(http.StatusOK, resp)
	})

	e.POST("/task/suite", TaskSuiteHandle)
	e.POST("/task/theme", TaskThemeHandle)

	e.GET("/themes", ThemesHandle)
	e.GET("/themes/:id", ThemeHandle)
	e.GET("/suites", SuitesHandle)
	e.GET("/suites/:suite_id", SuiteHandle)
	e.GET("/suites/:id/doread", SuiteReadHandle)
	e.GET("/suites/:id/dolike", SuiteLikeHandle)

	devopsGroup := e.Group("/devops")
	devopsGroup.POST("/initdb", InitDBHandle)
	devopsGroup.GET("", DevopsHandle)
	e.Static("/image/*filepath", Config.BasePath)

	addr := fmt.Sprintf("%s%s", Config.IP, Config.Port)
	fmt.Printf("serve: http://%s\n", addr)
	// server := http.Server{
	// 	Addr:    addr,
	// 	Handler: mux,
	// }
	// server.ListenAndServe()
	e.Logger.Fatal(e.Start(Config.Port))
}
