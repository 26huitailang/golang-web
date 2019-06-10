package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	// "log"
	"net/http"
	"os"
	"reflect"

	"github.com/26huitailang/golang-web/downloadsuite/suite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var config *Configuration
var DB *gorm.DB

type Configuration struct {
	BasePath    string `json:"base_path"`
	IP          string `json:"ip"`
	Port        string `json:"port"`
	DeployLevel int    `json:"deploy_level"`
}

// 初始化文件结构
func init() {
	var err error
	// log
	// logfile, err := os.OpenFile("main.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	// 	log.SetOutput(logfile)
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	initConfiguration()
	// DB 小心:= 覆盖了声明的全局变量
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Panicf("DB connect error: %s", err)
	}
	DB.LogMode(true)

	// 迁移
	DB.SingularTable(true) // 单数表名
	DB.AutoMigrate(&Theme{}, &Suite{}, &Image{})
	// sqlite 对alter table的支持有限，不支持rename column和remove column
	// err = DB.Model(&Image{}).DropColumn("IsRead").Error

	// 模板预加载
	ReloadTemplates()
}

// 加载默认配置config.json
func initConfiguration() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Error opening JSON file:", err)
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error reading JSON data:", err)
	}
	json.Unmarshal(jsonData, &config)
	initCustomConfig()
	log.Println("config:", config)
}

// 加载自定义配置，覆盖默认配置
func initCustomConfig() {
	// 文件是否存在
	file, err := os.Open("config_custom.json")
	if err != nil {
		log.Warn("config_custom.json not existed!\nUse default config.json\n")
		return
	}
	defer file.Close()
	// 读取json为map
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// var customConfig Configuration
	var customMap map[string]interface{}
	json.Unmarshal(data, &customMap)

	// 遍历config结构体，判断是否有覆盖内容
	t := reflect.TypeOf(config).Elem()
	v := reflect.ValueOf(config).Elem()
	for i := 0; i < t.NumField(); i++ {
		// 比较tag
		fieldInfo := t.Field(i)
		tag := fieldInfo.Tag.Get("json")
		if value, ok := customMap[tag]; ok {
			log.Printf("tag: [%s %v] replaced by [%v]", tag, v.Field(i).Interface(), value)
			v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
		}
	}
}

func themes(c echo.Context) (err error) {
	var themes []Theme
	DB.Order("name").Find(&themes)

	c.Render(http.StatusOK, "layout:themes", themes)
	return
}

type themeQuery struct {
	IsRead bool `query:"is_read"`
}

func theme(c echo.Context) (err error) {
	themeID, err := strconv.Atoi(c.Param("id"))
	queryParams := c.QueryParams()
	query := new(themeQuery)
	if err = c.Bind(query); err != nil {
		return
	}
	log.Debugf("is_read: %v %T", query.IsRead, query.IsRead)
	if err != nil {
		return c.Redirect(404, "/")
	}

	var theme Theme
	var suites []Suite
	DB.Where("id = ?", themeID).First(&theme)

	if _, ok := queryParams["is_read"]; ok {
		DB.Model(&theme).Where("is_read = ?", query.IsRead).Related(&suites).Order("name")
	} else {
		DB.Model(&theme).Related(&suites).Order("name")
	}

	log.Debugf("theme api suites[%s]: %v", theme.Name, suites)
	data := struct {
		Theme  Theme
		Suites []Suite
	}{
		theme,
		suites,
	}
	return c.Render(http.StatusOK, "layout:theme", data)
}

func suiteHandle(c echo.Context) (err error) {
	var sutieID int
	sutieID, err = strconv.Atoi(c.Param("suite_id"))
	if err != nil {
		return c.Redirect(http.StatusNotFound, "/")
	}
	log.Debugf("suiteID: %d", sutieID)
	var suite Suite
	var images []Image
	DB.Where("id = ?", sutieID).Find(&suite)
	log.Debugf("suite: %v", suite)
	DB.Model(&suite).Related(&images).Order("name")
	log.Debugf("images: %v", images)
	data := struct {
		Name   string
		Images []Image
	}{
		suite.Name,
		images,
	}
	return c.Render(200, "layout:suite", data)
}

func index(c echo.Context) (err error) {
	http.Redirect(c.Response(), c.Request(), "/themes", 301)
	return nil
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
		t := suite.NewTheme(form.URL, config.BasePath)
		t.DownloadOneTheme()
		fmt.Printf("%v", t)
	}()
	return c.String(http.StatusAccepted, "task theme sent ...")
}

func initDB(c echo.Context) (err error) {
	// todo websocket，异步？
	log.Println("droppig table ...")
	DB.DropTableIfExists(Theme{}, Suite{}, Image{})
	log.Println("migrating table ...")
	DB.AutoMigrate(Theme{}, Suite{}, Image{})
	log.Println("start init db ...")
	InitTheme(config)
	return c.String(200, "finish init db!\n")
}

// todo: /suite/:id/like 翻转操作，时间限制，3s一次

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

func main() {
	e := echo.New()
	if config.DeployLevel >= Development {
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
	e.GET("/", index)
	e.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")
		resp := fmt.Sprintf("Hello, %s!", name)
		return c.String(http.StatusOK, resp)
	})

	e.POST("/task/suite", taskSuite)
	e.POST("/task/theme", taskTheme)

	e.GET("/themes", themes)
	e.GET("/themes/:id", theme)
	e.GET("/themes/:theme_id/suites/:suite_id", suiteHandle)

	e.POST("/devops/initdb", initDB)
	e.Static("/image/*filepath", config.BasePath)

	addr := fmt.Sprintf("%s%s", config.IP, config.Port)
	fmt.Printf("serve: http://%s\n", addr)
	// server := http.Server{
	// 	Addr:    addr,
	// 	Handler: mux,
	// }
	// server.ListenAndServe()
	e.Logger.Fatal(e.Start(config.Port))
}
