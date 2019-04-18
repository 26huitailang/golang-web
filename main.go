package main

import (
	//"net/http"

	"flag"
	"io"
	"os"

	"github.com/26huitailang/golang-web/routers"
	"github.com/gin-gonic/gin"
)

var baseFolder string

func init() {
	flag.StringVar(&baseFolder, "folder", "/home/pi/Downloads/meituri/丽柜", "default: /home/pi/Downloads/meituri/丽柜")
}

// Constants 初始化一些常量的中间件
func Constants() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("base_folder", baseFolder)
		c.Next()
	}
}

func main() {
	flag.Parse()
	r := gin.Default()
	// log
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// middleware
	r.Use(Constants())
	// init router
	routers.SetupRouter(r, baseFolder)

	r.Run()
}
