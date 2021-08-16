package downloadsuite

import (
	"encoding/json"
	"fmt"
	"github.com/26huitailang/golang_web/config"
	"github.com/gosuri/uiprogress"
	"github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// todo: 下载人员标签下的所有 https://www.lanvshen.com/t/5500/

var _ ISuiteOperator = (*MeituriSuite)(nil) // check implement interface

// MeituriSuite struct 第一页的URL，suite的标题，第一个的HTML内容
type MeituriSuite struct {
	BaseFolderPath   string      `json:"base_folder_path"`
	FirstPage        string      `json:"first_page"`
	FirstHTMLContent string      `json:"first_html_content"`
	OrgName          string      `json:"org_name"`
	Title            string      `json:"title"`
	SuiteFolderPath  string      `json:"suite_folder_path"`
	PageMax          int         `json:"page_max"`
	countFanOut      int         `json:"-"`
	ChanPage         chan string `json:"-"`
	ChFailedImg      chan string `json:"-"` // 下载失败img放回
	Parser           IMTRParser  `json:"-"`
	Images           []*ImageInfo
}

type ImageInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Path string `json:"path"`
}

// 需要的操作的接口，方便测试时stub
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
		Images:         []*ImageInfo{},
	}
	suite.FirstHTMLContent = suite.Parser.PageContent(firstPage)
	suite.Title = suite.Parser.ParseTitle(suite.FirstHTMLContent)
	suite.PageMax = suite.Parser.FindSuitePageMax(suite.FirstHTMLContent)
	suite.getSuiteFolderPath()
	return suite
}

// collectImages to collect every page images to a channel
func (s *MeituriSuite) collectImages() {
	go s.GetPageURLs()

	var chImgs []<-chan string
	for i := 0; i < s.countFanOut; i++ {
		ch := s.GetImgURLs()
		chImgs = append(chImgs, ch)
	}
	// 回收多个channel的结果
	chImg := Merge(chImgs...)
	for imgURL := range chImg {
		name := getNameFromURL(imgURL)
		img := &ImageInfo{
			name,
			imgURL,
			s.SuiteFolderPath,
		}
		log.Debug("append image to s.Images:", img)
		s.Images = append(s.Images, img)
	}
}

// Download 实现接口方法，下载chImg channel中的URL
func (s *MeituriSuite) Download() {
	s.collectImages()

	// 文件夹检查
	// 根据folder是不是基础路径来判断是否从org获取真实名称，并加入到路径中
	isFolderExist := IsFileOrFolderExists(s.SuiteFolderPath)
	if !isFolderExist {
		fmt.Println("创建文件夹: ", s.SuiteFolderPath)
		err := os.MkdirAll(s.SuiteFolderPath, os.ModePerm)
		CheckError(err)
	}

	var chDownloads []<-chan string
	chImg := make(chan *ImageInfo, 20)
	go func() {
		for _, img := range s.Images {
			chImg <- img
		}
		close(chImg)
	}()
	for i := 0; i < s.countFanOut; i++ {
		ch := s.downloader(chImg)
		chDownloads = append(chDownloads, ch)
	}

	if config.Config.UIProgress.Show {
		uiprogress.Start()
		bar := uiprogress.AddBar(len(s.Images)).AppendCompleted().PrependElapsed()

		// 回收下载结果
		finish := Merge(chDownloads...)

		for _ = range finish {
			bar.Incr()
			//fmt.Println("finish: ", ret)
		}
		//uiprogress.Stop()
		//time.Sleep(time.Millisecond * 100)
	} else {
		finish := Merge(chDownloads...)
		for v := range finish {
			log.Debug("finish: ", v)
		}
	}
	//time.Sleep(time.Millisecond * 100)
}

// 下载
func (s *MeituriSuite) downloader(inCh <-chan *ImageInfo) chan string {
	finish := make(chan string)

	go func() {
		defer close(finish)
		for imageInfo := range inCh {
			name := path.Join(s.SuiteFolderPath, imageInfo.Name)
			if IsFileOrFolderExists(name) {
				fmt.Println("已存在: ", name)
				continue
			}
			content := getImageContent(imageInfo.URL)
			if ioutil.WriteFile(name, content, 0644) == nil {
				finish <- imageInfo.URL
			} else {
				fmt.Println("failed: ", imageInfo, " 放入chFailedImg")
				s.ChFailedImg <- imageInfo.URL
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
			log.Debug("write page to s.ChanPage:", s.FirstPage)
			s.ChanPage <- s.FirstPage
		default:
			pageURL := s.generatePageURL(i)
			log.Debug("write page to s.ChanPage:", pageURL)
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
					log.Info("no more url, ChangePage status:", ok, url)
					return
				}
				log.Debug("Try to get content:", url)
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

// 尝试从suite页获取机构名解决theme和suite使用不一致的问题
func (s *MeituriSuite) getSuiteFolderPath() string {
	if s.SuiteFolderPath != "" {
		return s.SuiteFolderPath
	}

	// theme 处理时不提取name，统一在这里获取
	orgName := s.GetOrgName(s.FirstHTMLContent)
	s.SuiteFolderPath = filepath.Join(s.BaseFolderPath, orgName, s.Title)
	return s.SuiteFolderPath
}

func (s *MeituriSuite) GetOrgName(content string) string {
	if s.OrgName != "" {
		return s.OrgName
	}
	return parseOrgName(content)
}

// Produce to produce img info to nsq
func (s *MeituriSuite) Produce(producer *nsq.Producer, topic string) error {
	s.collectImages()
	for img := range s.Images {

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
		title = strings.Trim(title, "\n")
		title = strings.TrimSpace(title)
		log.Println("title does not include </a>:", title)
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
	log.Println("title includes:", title)
	return title
}
