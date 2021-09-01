package middleware

import (
	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/config"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type CustomContext struct {
	echo.Context
	Session *model.SessionValue
}

func (c *CustomContext) SetConfig() {
	log.Infoln("setting config...")
	c.Set("config", config.Config)
	log.Infoln("finish set config!")
}
