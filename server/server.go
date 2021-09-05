package server

import (
	"fmt"
	"github.com/26huitailang/golang_web/config"
	"github.com/26huitailang/golang_web/constants"
	mymiddleware "github.com/26huitailang/golang_web/middleware"
	"github.com/26huitailang/golang_web/router"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"

	// "log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

type CustomValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

var validate *validator.Validate

func (cv *CustomValidator) Validate(i interface{}) error {
	//en := en2.New()
	if err := cv.validator.Struct(i); err != nil {
		log.Errorf("Validate error: %s", err.Error())
		errV, ok := err.(validator.ValidationErrors)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		validatedErrors := errV.Translate(cv.trans)
		errInfo := make([]string, 0)
		for _, v := range validatedErrors {
			errInfo = append(errInfo, v)
		}
		return echo.NewHTTPError(http.StatusBadRequest, errInfo)
	}
	return nil
}

func CSRFSkipper(c echo.Context) bool {
	log.Info(c.Path())
	return c.Path() == "/login" || c.Path() == "/logout"
}

func NewServer() *echo.Echo {
	Init()
	e := echo.New()

	// validator, trans
	zhTrans := zh.New()
	uni := ut.New(zhTrans, zhTrans)
	trans, ok := uni.GetTranslator("zh")
	if !ok {
		log.Fatalf("uni.GetTranslator(%s) filed", "zh")
	}
	validate = validator.New()
	zh_translations.RegisterDefaultTranslations(validate, trans)
	e.Validator = &CustomValidator{validator: validate, trans: trans}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &mymiddleware.CustomContext{c, nil}
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
		Skipper:      CSRFSkipper,
		TokenLookup:  "header:X-CSRFToken",
		ContextKey:   "csrf",
		CookieName:   "_csrf",
		CookieMaxAge: 86400,
	}))
	// e.Use(middleware.JWT([]byte("secret")))

	var EchoTemplate = &config.Template{}
	e.Renderer = EchoTemplate

	//DB := database.DB()
	//store := &views.DatabaseStore{DB: DB}
	//handler := &views.Handler{Store: store}
	router.Router(e)

	addr := fmt.Sprintf("%s%s", config.Config.IP, config.Config.Port)
	fmt.Printf("serve: http://%s\n", addr)

	return e
}
