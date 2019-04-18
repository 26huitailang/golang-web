package controllers

import (
	"fmt"
	"net/http"

	"github.com/26huitailang/golang-web/models"
	"github.com/gin-gonic/gin"
)

func MeituriListHandler(c *gin.Context) {
	baseFolder := c.MustGet("base_folder").(string)
	suiteList := models.GetFolderList(baseFolder)
	fmt.Println(suiteList[0].Folder)
	// c.JSON(200, suiteList)
	// ctx.TplName = "meituri/list.html"
	c.HTML(http.StatusOK, "meituriList", gin.H{
		"suiteList": suiteList,
	})
}

func MeituriDetailHandler(c *gin.Context) {
	baseFolder := c.MustGet("base_folder").(string)
	title := c.Param("title")
	suite := models.Suite{Folder: title}
	suite.GetFileList(baseFolder)
	// ctx.Data["suite"] = suite
	fmt.Println(suite)
	// ctx.TplName = "meituri/view.html"
	c.HTML(http.StatusOK, "meituriDetail", gin.H{
		"suite": suite,
	})
}
