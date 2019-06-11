package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"syscall"
	"time"

	"github.com/26huitailang/golang-web/database"

	// "log"
	"net/http"
	"os"

	"github.com/26huitailang/golang-web/downloadsuite/suite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func startChild1() {
	cmd := exec.Command("/bin/sh", "-c", "sleep 1000")
	time.AfterFunc(10*time.Second, func() {
		fmt.Println("PID1=", cmd.Process.Pid)
		syscall.Kill(-cmd.Process.Pid, syscall.SIGQUIT)
		fmt.Println("killed")
	})
	fmt.Println("begin run")
	cmd.Run()
}

func startChild2() {
	for index := 0; index < 10; index++ {
		time.Sleep(1 * time.Second)
		fmt.Println(index)
	}
}

func taskSuite(c echo.Context) (err error) {
	// go startChild1()
	// go startChild2()
	go func() {
		s := suite.NewSuite("https://www.meituri.com/a/26718/")
		suite.DonwloadSuite(s, 5, "/Users/26huitailang/Downloads/mzitu_go", s.Title)
	}()
	return c.String(http.StatusAccepted, "task suite sent ...")
}

func taskTheme(c echo.Context) (err error) {
	var form struct {
		URL string `json:"url"`
	}
	err = json.NewDecoder(c.Request().Body).Decode(&form)
	if err != nil {
		return c.String(500, err.Error())
	}
	log.Println(form)

	go func() {
		t := suite.NewTheme(form.URL, Config.BasePath)
		t.DownloadOneTheme()
		fmt.Printf("%v", t)
	}()
	return c.String(http.StatusAccepted, "task theme sent ...")
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

func main() {
	e := echo.New()
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
	e.Use(middleware.Logger())
	e.Use(middleware.CSRF())
	// e.Use(middleware.JWT([]byte("secret")))

	var EchoTemplate = &Template{}
	e.Renderer = EchoTemplate

	// profiling
	// mux = httpprof.WrapRouter(mux)
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.GET("/", IndexHandle)
	e.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")
		resp := fmt.Sprintf("Hello, %s!", name)
		return c.String(http.StatusOK, resp)
	})

	e.POST("/task/suite", taskSuite)
	e.POST("/task/theme", taskTheme)

	e.GET("/themes", ThemesHandle)
	e.GET("/themes/:id", ThemeHandle)
	e.GET("/suites", SuitesHandle)
	e.GET("/suites/:suite_id", SuiteHandle)
	e.GET("/suites/:id/doread", SuiteReadHandle)
	e.GET("/suites/:id/dolike", SuiteLikeHandle)

	e.POST("/devops/initdb", InitDBHandle)
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
