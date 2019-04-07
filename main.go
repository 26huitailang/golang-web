package main

import (
	_ "beelog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/meituri/img", beego.AppConfig.String("imagefolder"))
	beego.Run()
}
