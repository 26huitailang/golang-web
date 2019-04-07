package routers

import (
	"beelog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/meituri", &controllers.MeituriListController{})
	beego.Router("/meituri/:title", &controllers.MeituriDetailController{})
}
