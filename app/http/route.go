package http

import (
	"github.com/26huitailang/golang-web/app/http/module/demo"
	"github.com/26huitailang/yogo/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist")

	demo.Register(r)
}
