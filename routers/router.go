package routers

import (
	"net/http"

	"github.com/26huitailang/golang-web/controllers"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, baseFolder string) {
	setupRouter(r, baseFolder)
}

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	// todo: 这里有一个相对路径的坑，在main.go中调用没有问题，但是单元测试会找不到
	r.AddFromFiles("meituriList", "views/meituri/list.html", "views/layout.html", "views/header.html")
	r.AddFromFiles("meituriDetail", "views/meituri/view.html", "views/layout.html", "views/header.html")
	return r
}
func setupRouter(router *gin.Engine, baseFolder string) {
	router.HTMLRender = createMyRender()
	// router.LoadHTMLGlob("views/meituri/*")
	router.Static("/static", "./static")
	router.Static("/image", baseFolder)
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.GET("/", func(c *gin.Context) {
		// c.Request.URL.Path = "/meituri"
		// router.HandleContext(c)
		c.Redirect(http.StatusMovedPermanently, "/meituri")
	})
	router.GET("/meituri", controllers.MeituriListHandler)
	router.GET("/meituri/:title", controllers.MeituriDetailHandler)
}
