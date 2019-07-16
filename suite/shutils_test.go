package suite

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImgUrls(t *testing.T) {
	data, err := ioutil.ReadFile("content.html")
	CheckError(err)
	divContent := parseDivContent(string(data))
	if divContent == "" {
		t.Fatal("div content is empty")
	}
	imgUrls := parseImg(divContent)
	for i, val := range imgUrls {
		t.Log(i)
		t.Log(val)
	}
}

func TestNewSuite(t *testing.T) {
	// todo: 试试gomock工具
	suiteURL := "https://www.meituri.com/a/26718/"
	title := "黑丝亮皮连衣超短裙 [森萝财团] [BETA-038] 写真集"
	suite := NewSuite(suiteURL)
	assert.Equal(t, suite.Title, title, "Title should be the same.")
}
