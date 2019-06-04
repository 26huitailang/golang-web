package suite

import (
	"io/ioutil"
	"testing"
)

func TestParseSuites(t *testing.T) {
	data, err := ioutil.ReadFile("content_theme.html")
	checkError(err)
	ret := parseSuites(string(data))
	if len(ret) != 40 {
		t.Fatal("parseSuites test data should be len 40, got", len(ret))
	}
}

// func TestNewSuite(t *testing.T) {
// 	// todo: 试试gomock工具
// 	suiteURL := "https://www.meituri.com/a/26718/"
// 	title := "黑丝亮皮连衣超短裙 [森萝财团] [BETA-038] 写真集"
// 	suite := NewSuite(suiteURL)
// 	assert.Equal(t, suite.Title, title, "Title should be the same.")
// }
