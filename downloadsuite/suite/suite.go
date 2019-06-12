/*
Package suite to download one suite images.

	Now support meituri.
*/
package suite

import (
	"fmt"
	"os"
	"path"
)

// ISuiteOperator interface 定义了Suite的操作集合
type ISuiteOperator interface {
	GetPageURLs(chan string)
	GetImgURLs(chPage <-chan string, chFailedImg <-chan string) <-chan string
	Download(chImg <-chan string, chFailedImg chan string, folderPath string) <-chan string
	GetOrgURL() string
}

// DonwloadSuite to download one suite
// 解析所有page后fan-out，获取img链接，merge结果，再次fan-out给downloader
func DonwloadSuite(iSuite ISuiteOperator, countFanOut int, folderPath string, title string) {
	chPage := make(chan string)
	chFailedImg := make(chan string) // 下载失败img放回
	go iSuite.GetPageURLs(chPage)

	var chImgs []<-chan string
	for i := 0; i < countFanOut; i++ {
		ch := iSuite.GetImgURLs(chPage, chFailedImg)
		chImgs = append(chImgs, ch)
	}
	// 回收多个channel的结果
	chImg := merge(chImgs...)

	// 文件夹检查
	// 从org获取真实名称，并加入到路径中
	// todo: 如果有两个呢？是否只取第一个
	themeURL := iSuite.GetOrgURL()
	theme := NewTheme(themeURL, folderPath)
	suiteFolderPath := path.Join(folderPath, theme.Name, title)
	isFolderExist := IsFileOrFolderExists(suiteFolderPath)
	if !isFolderExist {
		fmt.Println("创建文件夹: ", suiteFolderPath)
		err := os.MkdirAll(suiteFolderPath, os.ModePerm)
		checkError(err)
	}

	var chDownloads []<-chan string
	for i := 0; i < countFanOut; i++ {
		ch := iSuite.Download(chImg, chFailedImg, suiteFolderPath)
		chDownloads = append(chDownloads, ch)
	}

	// 回收下载结果
	finish := merge(chDownloads...)

	for ret := range finish {
		fmt.Println("finish: ", ret)
	}
}
