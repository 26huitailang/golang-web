package suite

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
)

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
	}
	t.init(folderToSave)
	return t
}

func (t *Theme) init(folderToSave string) {
	t.FirstPageContent = getPageContent(t.FirstURL)
	t.parseName()
	t.Path = path.Join(folderToSave, t.Name)
	if ok := IsFileOrFolderExists(t.Path); !ok {
		os.Mkdir(t.Path, 0644)
	}
	t.parseMaxPage()
}

func (t *Theme) parseMaxPage() {
	re := regexp.MustCompile(`html">([0-9]+)</a> <a class="a1`)
	pageContentRegexp, _ := regexp.Compile(`html" >([0-9]+)</a>(\s)<a href="(.+?) class="next">`)
	tmp := pageContentRegexp.FindString(t.FirstPageContent)
	intRe, _ := regexp.Compile(`[0-9]+`)
	pageStr := intRe.FindString(tmp)
	pageMax, err := strconv.Atoi(pageStr)
	checkError(err)
}
func (t *Theme) parseName() {
	re := regexp.MustCompile(`<h1>(.+?)</h1>`)
	name := re.FindStringSubmatch(t.FirstPageContent)
	println(name[1])
	t.Name = name[1]
}

func (t *Theme) genPages() {

}
func (t *Theme) GetSuites() {
	// 放入channel
}

// DwonloadOneTheme to download a theme
// theme 文件夹确认，不存在建立
func (t *Theme) DownloadOneTheme() {
	fmt.Println("Theme")
	for s := range t.Suites {
		DonwloadSuite(s, 3, t.Path, s.Title)
	}
}
