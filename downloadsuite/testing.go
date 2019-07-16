package downloadsuite

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

// StubParse 用于替换实际的IMTRParser接口
type StubParser struct{}

func (s StubParser) FindSuitePageMax(content string) (pageMax int) {
	return findSuitePageMax(content)
}

func (s StubParser) PageContent(url string) string {
	content, err := ioutil.ReadFile(url)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func (s StubParser) ParseTitle(content string) (title string) {
	return parseTitle(content)
}

func (s StubParser) ParseOrgURL(content string) (URL string) {
	return parseOrgURL(content)
}

func AssertNoError(t *testing.T, err error) {
	assert.NoErrorf(t, err, "Expect no err, got %v", err)
}
