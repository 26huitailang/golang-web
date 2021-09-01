package views

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/database"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"

	"github.com/26huitailang/golang_web/app/service/downloadsuite"
	"github.com/26huitailang/golang_web/config"

	log "github.com/sirupsen/logrus"
)

type IThemeStore interface {
	GetThemes() []model.Theme
}

type Handler struct {
	IThemeStore
	Store *database.DatabaseStore
}

type ThemeStore struct {
}

func (s *ThemeStore) GetThemes() []model.Theme {
	var themes []model.Theme
	database.NewDatabaseStore().DB().Order("name").Find(&themes)
	return themes
}

func IndexHandle(c echo.Context) (err error) {
	http.Redirect(c.Response(), c.Request(), "/themes", 301)
	return nil
}

func (h *Handler) ThemesHandle(c echo.Context) (err error) {
	themes := h.GetThemes()
	err = c.Render(http.StatusOK, "layout:themes", themes)
	if err != nil {
		panic("render layout:themes error")
	}
	return
}

type themeQuery struct {
	IsRead bool `query:"is_read"`
}

func (h *Handler) ThemeHandle(c echo.Context) (err error) {
	themeID, err := strconv.Atoi(c.Param("id"))
	println(themeID, "lwjeofiwjefijwef")
	queryParams := c.QueryParams()
	query := new(themeQuery)
	if err = c.Bind(query); err != nil {
		fmt.Printf("erererer1: %v", err)
		log.Debugf("is_read: %v %T", query.IsRead, query.IsRead)
		fmt.Printf("erererer2: %v", err)
		return c.Redirect(404, "/")
	}

	var theme model.Theme
	var suites []model.Suite

	var DB = h.Store.DB()

	DB.Where("id = ?", themeID).First(&theme)

	if _, ok := queryParams["is_read"]; ok {
		DB.Model(&theme).Where("is_read = ?", query.IsRead).Related(&suites).Order("name")
	} else {
		DB.Model(&theme).Related(&suites).Order("name")
	}

	log.Debugf("theme api suites[%s]: %v", theme.Name, suites)

	var countRead, countUnread int
	DB.Model(&model.Suite{}).Where("theme_id = ? and is_read = true", themeID).Count(&countRead)
	DB.Model(&model.Suite{}).Where("theme_id = ? and is_read = false", themeID).Count(&countUnread)
	data := struct {
		Theme       model.Theme
		Suites      []model.Suite
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

func (h *Handler) SuiteHandle(c echo.Context) (err error) {
	var sutieID int
	sutieID, err = strconv.Atoi(c.Param("suite_id"))
	if err != nil {
		return c.Redirect(http.StatusNotFound, "/")
	}
	log.Debugf("suiteID: %d", sutieID)
	var suite model.Suite
	var images []model.Image

	DB := h.Store.DB()

	DB.Where("id = ?", sutieID).Find(&suite)
	log.Debugf("downloadsuite: %v", suite)
	DB.Model(&suite).Related(&images).Order("name")
	log.Debugf("images: %v", images)
	data := struct {
		Suite  model.Suite
		Images []model.Image
	}{
		suite,
		images,
	}
	return c.Render(200, "layout:suite", data)
}

func (h *Handler) SuiteReadHandle(c echo.Context) (err error) {
	var suiteID int
	suiteID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusNotFound, "/")
	}
	var suite model.Suite

	DB := h.Store.DB()

	DB.First(&suite, suiteID)
	suite.IsRead = !suite.IsRead
	DB.Save(&suite)
	url := fmt.Sprintf("/suites/%d#%s", suiteID, "action-read")
	log.Debugln("redirect:", url)
	return c.Redirect(302, url)
}

// todo: 时间限制，3s一次
func (h *Handler) SuiteLikeHandle(c echo.Context) (err error) {
	var suiteID int
	suiteID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusNotFound, "/")
	}
	var suite model.Suite

	DB := h.Store.DB()

	DB.First(&suite, suiteID)
	suite.IsLike = !suite.IsLike
	suite.IsRead = true
	DB.Save(&suite)
	url := fmt.Sprintf("/suites/%d#%s", suiteID, "action-like")
	log.Debugln("redirect:", url)
	return c.Redirect(302, url)
}

// InitDBHandle is view to init database depends on local files
func (h *Handler) InitDBHandle(c echo.Context) (err error) {
	DB := h.Store.DB()

	// todo websocket，异步？
	log.Println("droppig table ...")
	DB.DropTableIfExists(model.Theme{}, model.Suite{}, model.Image{})
	log.Println("migrating table ...")
	DB.AutoMigrate(model.Theme{}, model.Suite{}, model.Image{})
	log.Println("start init db ...")
	// todo: 这里因为用了session，所以在提交前也是看不到任何数据的
	go downloadsuite.InitTheme(config.Config)
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
		p, err := os.FindProcess(cmd.Process.Pid)
		if err != nil {
			log.Error(err)
		}
		p.Signal(syscall.SIGQUIT)
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
		operator := downloadsuite.NewMeituriSuite(url, config.Config.MediaPath, downloadsuite.MeituriParser{})
		suite := downloadsuite.NewSuite(operator)
		suite.Download()
		// 重新加载进去
		downloadsuite.InitTheme(config.Config)
	}()
	return c.String(http.StatusAccepted, "task downloadsuite sent ...")
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
		t := downloadsuite.NewTheme(url, config.Config.MediaPath)
		t.DownloadOneTheme()
		fmt.Printf("%v", t)
		downloadsuite.InitTheme(config.Config)
	}()
	return c.String(http.StatusAccepted, "task theme sent ...")
}
