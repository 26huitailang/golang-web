package suite

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// MeituriSuite struct 第一页的URL，suite的标题，第一个的HTML内容
type MeituriSuite struct {
	firstPage       string
	Title           string
	firtHTMLContent string
	OrgURL          string
}

// NewSuite 初始化一个MeituriSuite结构
func NewSuite(firstPage string) *MeituriSuite {
	suite := &MeituriSuite{
		firstPage: firstPage,
	}
	suite.firtHTMLContent = getPageContent(firstPage)
	suite.parseTitle()
	suite.getOrgURL()
	return suite
}

// GetPageURLs 接口方法，生成每页的URL
func (suite *MeituriSuite) GetPageURLs(chPage chan string) {
	defer close(chPage)
	pageMax := findSuitePageMax(suite.firtHTMLContent)
	// 没有分页，返回firstPage即可
	for i := 1; i <= pageMax; i++ {
		switch i {
		case 1: // 第一页特殊
			chPage <- suite.firstPage
		default:
			pageURL := suite.generatePageURL(i)
			chPage <- pageURL
		}
	}
}

// GetImgURLs 实现接口方法，获取每页的ImgURL放入channel
func (suite *MeituriSuite) GetImgURLs(chPage <-chan string, chFailedImg <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case url, ok := <-chPage:
				if !ok {
					fmt.Println("no more url")
					return
				}
				content := getPageContent(url)
				divContent := parseDivContent(content)
				imgSrcs := parseImg(divContent)
				// fmt.Println(imgSrcs)
				for _, imgSrc := range imgSrcs {
					out <- imgSrc
				}
			case url := <-chFailedImg:
				out <- url
			}
		}
	}()
	return out
}

// Download 实现接口方法，下载chImg channel中的URL
func (suite *MeituriSuite) Download(chImg <-chan string, chFailedImg chan string, folderPath string) <-chan string {
	finish := make(chan string)
	go func() {
		downloader(chImg, finish, chFailedImg, folderPath)
	}()
	return finish
}

// 获取最大页码
func findSuitePageMax(firstHTMLContent string) (pageMax int) {
	pageContentRegexp, _ := regexp.Compile(`html">([0-9]+)</a> <a class="a1`)
	tmp := pageContentRegexp.FindString(firstHTMLContent)
	intRe, _ := regexp.Compile(`[0-9]+`)
	pageStr := intRe.FindString(tmp)
	pageMax, err := strconv.Atoi(pageStr)
	// 没有分页组件，表示就一页
	if err != nil {
		pageMax = 1
	}
	return
}

// 根绝页码和firstPage构建其余的page的URL
func (suite *MeituriSuite) generatePageURL(page int) string {
	// https://www.meituri.com/a/26718/2.html
	pageStr := strconv.Itoa(page)
	return suite.firstPage + pageStr + ".html"
}

// todo: 如果下面的方法是MeituriSuite使用并不能公共使用的话，写到struct下面
// 下载，消费者
func downloader(inCh <-chan string, finishCh chan string, chFailedImg chan string, suiteFolderPath string) {
	for url := range inCh {
		nameStrings := strings.Split(url, "/")
		name := nameStrings[len(nameStrings)-1]
		name = path.Join(suiteFolderPath, name)
		if IsFileOrFolderExists(name) {
			fmt.Println("已存在: ", name)
			continue
		}
		content := getImageContent(url)
		if ioutil.WriteFile(name, content, 0644) == nil {
			finishCh <- url
		} else {
			fmt.Println("failed: ", url, " 放入chFailedImg")
			chFailedImg <- url
		}
	}
	defer close(finishCh)
}

// 解析div class="content" 部分
func parseDivContent(content string) string {
	divContentRegexp, _ := regexp.Compile(`<div class="content">([\s\S]*?)</div>`)
	divContent := divContentRegexp.FindString(content)
	return divContent
}

// 解析所有的img src
func parseImg(content string) []string {
	var imgSrcs []string
	imgSrcRegexp, _ := regexp.Compile(`img src="(.+?)" alt`)
	imgSrcsSlice := imgSrcRegexp.FindAllStringSubmatch(content, -1)
	for _, valSlice := range imgSrcsSlice {
		imgSrcs = append(imgSrcs, valSlice[len(valSlice)-1])
	}
	return imgSrcs
}

// 从URL获取文件内容
func getImageContent(url string) []byte {
	body := getURLContent(url)
	return body
}

func (suite *MeituriSuite) parseTitle() {
	titleRegexp, _ := regexp.Compile(`<h1>(.+?)</h1>`)
	title := titleRegexp.FindStringSubmatch(suite.firtHTMLContent)
	suite.Title = title[1]
}

func (suite *MeituriSuite) getOrgURL() {
	url := parseOrgURL(suite.firtHTMLContent)
	suite.OrgURL = url
}

func parseOrgURL(content string) (url string) {
	// println(content)
	re := regexp.MustCompile(`<p>拍摄机构：([\s\S]*?)<a href="(.*?)" target="_blank">`) // 非贪婪
	texts := re.FindStringSubmatch(content)
	println("texts:", texts)
	url = texts[2]
	println("parseOrgURL:", url)
	return
}

// GetOrgURL 获取对象的该属性
func (suite *MeituriSuite) GetOrgURL() string {
	return suite.OrgURL
}
