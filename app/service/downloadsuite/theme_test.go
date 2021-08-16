package downloadsuite

import (
	"io/ioutil"
	"testing"
)

func TestParseSuites(t *testing.T) {
	data, err := ioutil.ReadFile("mtr_theme.html")
	CheckError(err)
	ret := parseSuites(string(data))
	if len(ret) != 40 {
		t.Fatal("parseSuites test data should be len 40, got", len(ret))
	}
}

// func TestNewSuite(t *testing.T) {
// 	// todo: 试试gomock工具
// 	suiteURL := "https://www.meituri.com/a/26718/"
// 	title := "黑丝亮皮连衣超短裙 [森萝财团] [BETA-038] 写真集"
// 	downloadsuite := NewMeituriSuite(suiteURL)
// 	assert.Equal(t, downloadsuite.Title, title, "Title should be the same.")
// }

func TestParseThemePageMax(t *testing.T) {
	input := []string{
		`<center><div id="pages" class="text-c">
		第一页 <a href="javascript:void(0)" class="current">1</a>
		 <a href="/x/82/index_1.html" >2</a>
		 <a href="/x/82/index_2.html" >3</a>
		 <a href="/x/82/index_3.html" >4</a>
		<a href="/x/82/index_1.html" class="next">下一页</a></div></center>`,
		"",
	}
	output := []int{4, 1}
	for i, content := range input {
		out := parseThemeMaxPage(content)
		if out != output[i] {
			t.Fatalf("Expected %d, Got %d", output[i], out)
		}
	}
}
