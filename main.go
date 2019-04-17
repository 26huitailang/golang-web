package main

import (
	//"net/http"

	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

type Image struct {
	Title string
	Path  string
}
type Suite struct {
	Images []*Image
	Folder string
}

var baseFolder string

func init() {
	flag.StringVar(&baseFolder, "folder", "/home/pi/Downloads/meituri/丽柜", "default: /home/pi/Downloads/meituri/丽柜")
}

func getFolderList(path string) []*Suite {
	fmt.Println(baseFolder)
	suiteList := make([]*Suite, 0)
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			suite := &Suite{Folder: file.Name()}
			suiteList = append(suiteList, suite)
		} else {
			continue
		}
	}
	return suiteList
}
func getFileList(suite *Suite) *Suite {
	imageList := make([]*Image, 0)
	suitePath := filepath.Join(baseFolder, suite.Folder)
	files, _ := ioutil.ReadDir(suitePath)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), "jpg") {
			imagePath := filepath.Join(suite.Folder, file.Name())
			image := &Image{Title: file.Name(), Path: imagePath}
			imageList = append(imageList, image)
		}
	}
	suite.Images = imageList
	return suite
}

func MeituriListHandler(c *gin.Context) {
	suiteList := getFolderList(baseFolder)
	fmt.Println(suiteList[0].Folder)
	// c.JSON(200, suiteList)
	// ctx.TplName = "meituri/list.html"
	c.HTML(http.StatusOK, "meituriList", gin.H{
		"suiteList": suiteList,
	})
}

func MeituriDetailHandler(c *gin.Context) {
	title := c.Param("title")
	suite := Suite{Folder: title}
	getFileList(&suite)
	// ctx.Data["suite"] = suite
	fmt.Println(suite)
	// ctx.TplName = "meituri/view.html"
	c.HTML(http.StatusOK, "meituriDetail", gin.H{
		"suite": suite,
	})
}

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("meituriList", "views/meituri/list.html", "views/layout.html", "views/header.html")
	r.AddFromFiles("meituriDetail", "views/meituri/view.html", "views/layout.html", "views/header.html")
	return r
}

func setupRouter() *gin.Engine {
	router := gin.Default()
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
	router.GET("/meituri", MeituriListHandler)
	router.GET("/meituri/:title", MeituriDetailHandler)
	return router
}

func main() {
	flag.Parse()
	// init router
	router := setupRouter()
	router.Run()
}
