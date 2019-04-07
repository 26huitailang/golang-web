package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type MainController struct {
	beego.Controller
}
type MeituriListController struct {
	beego.Controller
}
type MeituriDetailController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Redirect("/meituri", 302)
}

type Image struct {
	Title string
	Path  string
}
type Suite struct {
	Images []*Image
	Folder string
}

var baseFolder = beego.AppConfig.String("imagefolder")

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
			imagePath := filepath.Join("img", suite.Folder, file.Name())
			image := &Image{Title: file.Name(), Path: imagePath}
			imageList = append(imageList, image)
		}
	}
	suite.Images = imageList
	return suite
}

//func init() {
//	userCur, err := user.Current()
//	if err != nil {
//		panic("No user home path")
//	}
//	userHome = userCur.HomeDir
//	baseFolder = filepath.Join(userHome, "Downloads/meituri/丽柜")
//}
func (ctx *MeituriListController) Get() {
	ctx.Data["a"] = "h"
	suiteList := getFolderList(baseFolder)
	fmt.Println(suiteList)
	ctx.Data["suiteList"] = suiteList
	ctx.TplName = "meituri/list.html"
}
func (ctx *MeituriDetailController) Get() {
	title := ctx.Ctx.Input.Param(":title")
	suite := Suite{Folder: title}
	getFileList(&suite)
	ctx.Data["suite"] = suite
	fmt.Println(suite)
	ctx.TplName = "meituri/view.html"
}
