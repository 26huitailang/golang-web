package main

import (
	"net/http"
	"strconv"

	"github.com/26huitailang/golang-web/models"
	"github.com/labstack/echo"

	// "github.com/labstack/gommon/log"
	log "github.com/sirupsen/logrus"
)

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
	data := struct {
		Theme  models.Theme
		Suites []models.Suite
	}{
		theme,
		suites,
	}
	return c.Render(http.StatusOK, "layout:theme", data)
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
		Name   string
		Images []models.Image
	}{
		suite.Name,
		images,
	}
	return c.Render(200, "layout:suite", data)
}

// todo: /suite/:id/like 翻转操作，时间限制，3s一次

// InitDBHandle is view to init database depends on local files
func InitDBHandle(c echo.Context) (err error) {
	// todo websocket，异步？
	log.Println("droppig table ...")
	DB.DropTableIfExists(models.Theme{}, models.Suite{}, models.Image{})
	log.Println("migrating table ...")
	DB.AutoMigrate(models.Theme{}, models.Suite{}, models.Image{})
	log.Println("start init db ...")
	// var config = config.Config
	// config.InitTheme()
	return c.String(200, "finish init db!\n")
}
