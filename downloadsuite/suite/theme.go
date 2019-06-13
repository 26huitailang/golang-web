package suite

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
)

// Theme 对应meituri机构
type Theme struct {
	FirstURL         string
	Name             string
	Path             string
	FirstPageContent string
	MaxPage          int
	Pages            chan string
	Suites           chan *MeituriSuite
}

func NewTheme(firstPage, folderToSave string) *Theme {
	// https://www.meituri.com/x/82/
	// https://www.meituri.com/x/82/index_1.html
	t := &Theme{
		FirstURL: firstPage,
		Pages:    make(chan string),
		Suites:   make(chan *MeituriSuite),
	}
	t.init(folderToSave)
	return t
}

func (t *Theme) init(folderToSave string) {
	t.FirstPageContent = getPageContent(t.FirstURL)
	t.parseName()
	t.Path = path.Join(folderToSave, t.Name)
	if ok := IsFileOrFolderExists(t.Path); !ok {
		os.Mkdir(t.Path, 0700)
	}
	t.getMaxPage()
}

func (t *Theme) getMaxPage() {
	t.MaxPage = parseThemeMaxPage(t.FirstPageContent)
}

func parseThemeMaxPage(FirstPageContent string) (pageMax int) {
	re := regexp.MustCompile(`html" >([0-9]+)</a>(\s*)<a href="(.+?) class="next">`)
	tmp := re.FindString(FirstPageContent)
	intRe := regexp.MustCompile(`[0-9]+`)
	pageStr := intRe.FindString(tmp)
	pageMax, err := strconv.Atoi(pageStr)
	// 单页，下面没有翻页组件
	if err != nil {
		pageMax = 1
	}
	return
}

func (t *Theme) parseName() {
	re := regexp.MustCompile(`<h1>(.+?)</h1>`)
	name := re.FindStringSubmatch(t.FirstPageContent)
	t.Name = name[1]
}

func (t *Theme) String() string {
	return fmt.Sprintf("Name: %s | Page: %d", t.Name, t.MaxPage)
}

func (t *Theme) genPages() {
	// https://www.meituri.com/x/82/
	// https://www.meituri.com/x/82/index_1.html
	for index := 0; index < t.MaxPage; index++ {
		switch index {
		case 0:
			t.Pages <- t.FirstURL
		default:
			URL := fmt.Sprintf("%sindex_%d.html", t.FirstURL, index)
			t.Pages <- URL
		}
	}
	close(t.Pages)
}

func parseSuites(pageContent string) (suiteURLs []string) {
	// log.Println("pageContent:", pageContent)
	reHezi := regexp.MustCompile(`<div class="hezi">([\d\D]*)</div>`)
	tmp := reHezi.FindString(pageContent)
	// log.Println("tmp:", tmp)
	reSuiteHref := regexp.MustCompile(`(.*?)a href="(.*?)" target="_blank"><img(.*?)`)
	groups := reSuiteHref.FindAllStringSubmatch(tmp, -1)
	for _, group := range groups {
		suiteURLs = append(suiteURLs, group[2])
	}
	// log.Println("suiteURLs: ", suiteURLs)
	return
}

func (t *Theme) genSuites() {
	// 放入channel
	for pageURL := range t.Pages {
		log.Println("Page:", pageURL)
		pageContent := getPageContent(pageURL)
		suiteURLs := parseSuites(pageContent)
		for _, suiteURL := range suiteURLs {
			suite := NewSuite(suiteURL)
			t.Suites <- suite
			log.Println(suiteURL)
		}
	}
	close(t.Suites)
}

// DownloadOneTheme to download a theme
// theme 文件夹确认，不存在建立
func (t *Theme) DownloadOneTheme() {
	fmt.Println("Theme")
	go t.genPages()
	go t.genSuites()
	for s := range t.Suites {
		DonwloadSuite(s, 3, t.Path, s.Title, true)
	}
	log.Println("DownloadOneTheme finished!")
}
