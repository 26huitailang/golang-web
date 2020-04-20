package downloadsuite

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var _ ISuiteOperator = (*MeituriSuite)(nil) // check implement interface

// MeituriSuite struct 第一页的URL，suite的标题，第一个的HTML内容
type MeituriSuite struct {
	BaseFolderPath   string      `json:"base_folder_path"`
	FirstPage        string      `json:"first_page"`
	FirstHTMLContent string      `json:"first_html_content"`
	OrgURL           string      `json:"org_url"`
	Title            string      `json:"title"`
	SuiteFolderPath  string      `json:"suite_folder_path"`
	PageMax          int         `json:"page_max"`
	countFanOut      int         `json:"-"`
	ChanPage         chan string `json:"-"`
	ChFailedImg      chan string `json:"-"` // 下载失败img放回
	Parser           IMTRParser  `json:"-"`
}

type ImageInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Path string `json:"path"`
}

type IMTRParser interface {
	PageContent(url string) string
	ParseTitle(content string) (title string)
	ParseOrgURL(content string) (URL string)
	FindSuitePageMax(content string) (pageMax int)
}

type MeituriParser struct{}

// NewMeituriSuite 初始化一个MeituriSuite结构
func NewMeituriSuite(firstPage string, folderPath string, parser IMTRParser) *MeituriSuite {
	suite := &MeituriSuite{
		FirstPage:      firstPage,
		countFanOut:    5,
		BaseFolderPath: folderPath,
		ChanPage:       make(chan string),
		ChFailedImg:    make(chan string),
		Parser:         parser,
	}
	suite.FirstHTMLContent = suite.Parser.PageContent(firstPage)
	suite.Title = suite.Parser.ParseTitle(suite.FirstHTMLContent)
	suite.OrgURL = suite.Parser.ParseOrgURL(suite.FirstHTMLContent)
	suite.PageMax = suite.Parser.FindSuitePageMax(suite.FirstHTMLContent)
	suite.getSuiteFolderPath()
	return suite
}

// collectImages to collect every page images to a channel
func (s *MeituriSuite) collectImages() <-chan string {
	go s.GetPageURLs()

	var chImgs []<-chan string
	for i := 0; i < s.countFanOut; i++ {
		ch := s.GetImgURLs()
		chImgs = append(chImgs, ch)
	}
	// 回收多个channel的结果
	chImg := Merge(chImgs...)
	return chImg
}

// Download 实现接口方法，下载chImg channel中的URL
func (s *MeituriSuite) Download() {
	chImg := s.collectImages()

	// 文件夹检查
	// 根据folder是不是基础路径来判断是否从org获取真实名称，并加入到路径中
	isFolderExist := IsFileOrFolderExists(s.SuiteFolderPath)
	if !isFolderExist {
		fmt.Println("创建文件夹: ", s.SuiteFolderPath)
		err := os.MkdirAll(s.SuiteFolderPath, os.ModePerm)
		CheckError(err)
	}

	var chDownloads []<-chan string
	for i := 0; i < s.countFanOut; i++ {
		ch := s.downloader(chImg)
		chDownloads = append(chDownloads, ch)
	}

	// 回收下载结果
	finish := Merge(chDownloads...)

	for ret := range finish {
		fmt.Println("finish: ", ret)
	}
}

// todo: 如果下面的方法是MeituriSuite使用并不能公共使用的话，写到struct下面
// 下载，消费者
func (s *MeituriSuite) downloader(inCh <-chan string) chan string {
	finish := make(chan string)

	go func() {
		defer close(finish)
		for url := range inCh {
			name := getNameFromURL(url)
			name = path.Join(s.SuiteFolderPath, name)
			if IsFileOrFolderExists(name) {
				fmt.Println("已存在: ", name)
				continue
			}
			content := getImageContent(url)
			if ioutil.WriteFile(name, content, 0644) == nil {
				finish <- url
			} else {
				fmt.Println("failed: ", url, " 放入chFailedImg")
				s.ChFailedImg <- url
			}
		}
	}()
	return finish
}

// GetPageURLs 接口方法，生成每页的URL
func (s *MeituriSuite) GetPageURLs() {
	getPageURLs(s)
}

func getPageURLs(s *MeituriSuite) {
	defer close(s.ChanPage)
	// 没有分页，返回firstPage即可
	for i := 1; i <= s.PageMax; i++ {
		switch i {
		case 1: // 第一页特殊
			s.ChanPage <- s.FirstPage
		default:
			pageURL := s.generatePageURL(i)
			s.ChanPage <- pageURL
		}
	}
}

// GetImgURLs 实现接口方法，获取每页的ImgURL放入channel
func (s *MeituriSuite) GetImgURLs() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case url, ok := <-s.ChanPage:
				if !ok {
					fmt.Println("no more url")
					return
				}
				content := GetPageContent(url)
				divContent := parseDivContent(content)
				imgSrcs := parseImg(divContent)
				// fmt.Println(imgSrcs)
				for _, imgSrc := range imgSrcs {
					out <- imgSrc
				}
			case url := <-s.ChFailedImg:
				out <- url
			}
		}
	}()
	return out
}

// 根绝页码和firstPage构建其余的page的URL
func (s *MeituriSuite) generatePageURL(page int) string {
	// https://www.meituri.com/a/26718/2.html
	pageStr := strconv.Itoa(page)
	return s.FirstPage + pageStr + ".html"
}

func getNameFromURL(url string) string {
	nameStrings := strings.Split(url, "/")
	name := nameStrings[len(nameStrings)-1]
	return name
}

func (s *MeituriSuite) getSuiteFolderPath() string {
	if s.SuiteFolderPath != "" {
		return s.SuiteFolderPath
	}

	themeURL := s.OrgURL
	theme := NewTheme(themeURL, s.BaseFolderPath)
	orgName := theme.Name
	//orgName := parseOrgName(s.FirstHTMLContent)
	isIncluded := strings.Contains(s.BaseFolderPath, orgName)
	if isIncluded {
		s.SuiteFolderPath = path.Join(s.BaseFolderPath, s.Title)
	} else {
		s.SuiteFolderPath = path.Join(s.BaseFolderPath, orgName, s.Title)
	}
	return s.SuiteFolderPath
}

// Produce to produce img info to nsq
func (s *MeituriSuite) Produce(producer *nsq.Producer, topic string) error {
	chImages := s.collectImages()
	for imgURL := range chImages {
		name := getNameFromURL(imgURL)
		img := &ImageInfo{
			name,
			imgURL,
			s.SuiteFolderPath,
		}
		data, err := json.Marshal(img)
		if err != nil {
			return errors.Errorf("marshal image error: %s", err.Error())
		}
		err = producer.Publish(topic, data)
		if err != nil {
			return errors.Errorf("suite produce image error: %s", err.Error())
		}
	}
	return nil
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
	body := GetURLContent(url)
	return body
}

func (p MeituriParser) ParseTitle(firtHTMLContent string) (title string) {
	title = parseTitle(firtHTMLContent)
	return
}

func (p MeituriParser) ParseOrgURL(content string) (URL string) {
	URL = parseOrgURL(content)
	return
}

func (p MeituriParser) PageContent(url string) string {
	return GetPageContent(url)
}

// 获取最大页码
func (p MeituriParser) FindSuitePageMax(firstHTMLContent string) (pageMax int) {
	return findSuitePageMax(firstHTMLContent)
}

func findSuitePageMax(content string) int {
	pageContentRegexp, _ := regexp.Compile(`html">([0-9]+)</a> <a class="a1`)
	tmp := pageContentRegexp.FindString(content)
	intRe, _ := regexp.Compile(`[0-9]+`)
	pageStr := intRe.FindString(tmp)
	pageMax, err := strconv.Atoi(pageStr)
	// 没有分页组件，表示就一页
	if err != nil {
		pageMax = 1
	}
	return pageMax
}

func parseTitle(content string) string {
	titleRegexp, _ := regexp.Compile(`<h1>(.+?)</h1>`)
	rets := titleRegexp.FindStringSubmatch(content)
	return rets[1]
}

func parseOrgURL(content string) string {
	re := regexp.MustCompile(`<p>拍摄机构：([\s\S]*?)<a href="(.*?)" target="_blank">`) // 非贪婪
	texts := re.FindStringSubmatch(content)
	URL := texts[2]
	//println("ParseOrgURL:", URL)
	return URL
}

func parseOrgName(content string) string {
	re := regexp.MustCompile(`<p>拍摄机构：([\s\S]*?)</p>`)
	texts := re.FindStringSubmatch(content)
	title := texts[1]

	// 标题无链接
	/*
		<p>拍摄机构：SIW斯文传媒</p>
	*/
	if !strings.Contains(title, "</a>") {
		return title
	}

	// 有链接的情况，取第一个
	/*
		<p>拍摄机构：
			<a href="https://www.lanvshen.com/x/12/" target="_blank">异思趣向</a>
			<a href="https://www.lanvshen.com/x/13/" target="_blank">丝足便当</a>
		</p>
	*/
	re = regexp.MustCompile(`<a href="(.*?)" target="_blank">(.*?)</a>`) // 非贪婪
	texts = re.FindStringSubmatch(content)
	title = texts[2]
	return title
}
