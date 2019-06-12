package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	"github.com/labstack/echo/middleware"

	"github.com/26huitailang/golang-web/downloadsuite/suite"
	"github.com/26huitailang/golang-web/models"
	"github.com/labstack/echo"

	log "github.com/sirupsen/logrus"
)

func IndexHandle(c echo.Context) (err error) {
	http.Redirect(c.Response(), c.Request(), "/themes", 301)
	return nil
}

func ThemesHandle(c echo.Context) (err error) {
	var themes []models.Theme
	DB.Order("name").Find(&themes)

	c.Render(http.StatusOK, "layout:themes", themes)
	return
}

type themeQuery struct {
	IsRead bool `query:"is_read"`
}

func ThemeHandle(c echo.Context) (err error) {
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

	var theme models.Theme
	var suites []models.Suite
	DB.Where("id = ?", themeID).First(&theme)

	if _, ok := queryParams["is_read"]; ok {
		DB.Model(&theme).Where("is_read = ?", query.IsRead).Related(&suites).Order("name")
	} else {
		DB.Model(&theme).Related(&suites).Order("name")
	}

	log.Debugf("theme api suites[%s]: %v", theme.Name, suites)

	var countRead, countUnread int
	DB.Model(&models.Suite{}).Where("theme_id = ? and is_read = true", themeID).Count(&countRead)
	DB.Model(&models.Suite{}).Where("theme_id = ? and is_read = false", themeID).Count(&countUnread)
	data := struct {
		Theme       models.Theme
		Suites      []models.Suite
		CountRead   int
		CountUnread int
	}{
		theme,
		suites,
		countRead,
		countUnread,
	}
	return c.Render(http.StatusOK, "layout:theme", data)
}

type suitesQuery struct {
	IsLike bool `query:"is_like"`
}

func SuitesHandle(c echo.Context) (err error) {
	query := new(suitesQuery)
	if err = c.Bind(query); err != nil {
		return
	}
	var suites []models.Suite
	DB.Where("is_like = ?", query.IsLike).Find(&suites)
	return c.Render(200, "layout:suites", suites)
}

func SuiteHandle(c echo.Context) (err error) {
	var sutieID int
	sutieID, err = strconv.Atoi(c.Param("suite_id"))
	if err != nil {
		return c.Redirect(http.StatusNotFound, "/")
	}
	log.Debugf("suiteID: %d", sutieID)
	var suite models.Suite
	var images []models.Image
	DB.Where("id = ?", sutieID).Find(&suite)
	log.Debugf("suite: %v", suite)
	DB.Model(&suite).Related(&images).Order("name")
	log.Debugf("images: %v", images)
	data := struct {
		Suite  models.Suite
		Images []models.Image
	}{
		suite,
		images,
	}
	return c.Render(200, "layout:suite", data)
}

func SuiteReadHandle(c echo.Context) (err error) {
	var suiteID int
	suiteID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusNotFound, "/")
	}
	var suite models.Suite
	DB.First(&suite, suiteID)
	suite.IsRead = !suite.IsRead
	DB.Save(&suite)
	url := fmt.Sprintf("/suites/%d#%s", suiteID, "action-read")
	log.Debugln("redirect:", url)
	return c.Redirect(302, url)
}

// todo: 时间限制，3s一次
func SuiteLikeHandle(c echo.Context) (err error) {
	var suiteID int
	suiteID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusNotFound, "/")
	}
	var suite models.Suite
	DB.First(&suite, suiteID)
	suite.IsLike = !suite.IsLike
	suite.IsRead = true
	DB.Save(&suite)
	url := fmt.Sprintf("/suites/%d#%s", suiteID, "action-like")
	log.Debugln("redirect:", url)
	return c.Redirect(302, url)
}

// InitDBHandle is view to init database depends on local files
func InitDBHandle(c echo.Context) (err error) {
	// todo websocket，异步？
	log.Println("droppig table ...")
	DB.DropTableIfExists(models.Theme{}, models.Suite{}, models.Image{})
	log.Println("migrating table ...")
	DB.AutoMigrate(models.Theme{}, models.Suite{}, models.Image{})
	log.Println("start init db ...")
	// todo: 这里因为用了session，所以在提交前也是看不到任何数据的
	go Config.InitTheme()
	return c.Redirect(302, "/")
}

func DevopsHandle(c echo.Context) (err error) {
	csrf := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	return c.Render(200, "layout:devops", csrf)
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

func TaskSuiteHandle(c echo.Context) (err error) {
	// go startChild1()
	// go startChild2()
	url := c.FormValue("url")
	log.Infoln("url:", url)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("%v", err)
			}
		}()
		s := suite.NewSuite(url)
		suite.DonwloadSuite(s, 5, Config.BasePath, s.Title, false)
		// 重新加载进去
		Config.InitTheme()
	}()
	return c.String(http.StatusAccepted, "task suite sent ...")
}

func TaskThemeHandle(c echo.Context) (err error) {

	// var form struct {
	// 	URL string `json:"url"`
	// }
	url := c.FormValue("url")
	// err = json.NewDecoder(c.Request().Body).Decode(&form)
	// log.Errorf("%v", err)
	// if err != nil {
	// 	return c.String(500, err.Error())
	// }
	log.Println("url:", url)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("%v", err)
			}
		}()
		t := suite.NewTheme(url, Config.BasePath)
		t.DownloadOneTheme()
		fmt.Printf("%v", t)
		Config.InitTheme()
	}()
	return c.String(http.StatusAccepted, "task theme sent ...")
}
