package downloadsuite

import (
	"io/ioutil"
	"testing"
)

func TestGetImgUrls(t *testing.T) {
	data, err := ioutil.ReadFile("mtr_suite.html")
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
