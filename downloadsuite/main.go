package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/26huitailang/golang-web/downloadsuite/suite"
)

// todo: 接收参数，只承担一个下载的角色
// 下载URL/保存的位置

var firstPage string
var folderSave string
var isTheme bool

func init() {
	// https://www.meituri.com/a/26718/
	flag.StringVar(&firstPage, "url", "", "Suite首页")
	flag.StringVar(&folderSave, "folder", "", "保存的suite路径")
	flag.BoolVar(&isTheme, "t", false, "Theme标记")
}

func main() {
	flag.Parse()
	if firstPage == "" {
		panic("page is empty")
	} else if folderSave == "" {
		panic("folder is empty")
	} else if isAbs := path.IsAbs(folderSave); !isAbs {
		panic("folder should be absolute")
	}
	if isTheme {
		t := suite.NewTheme(firstPage, folderSave)
		t.DownloadOneTheme()
	} else {
		s := suite.NewSuite(firstPage)
		suite.DonwloadSuite(s, 5, folderSave, s.Title, isTheme)
	}
	fmt.Println("✅")
}
