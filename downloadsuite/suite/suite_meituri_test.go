package suite

import (
	"io/ioutil"
	"testing"
)

func TestParseOrgURL(t *testing.T) {
	target := "https://www.meituri.com/x/82/"
	data, err := ioutil.ReadFile("content.html")
	checkError(err)
	ret := parseOrgURL(string(data))
	if ret != target {
		t.Fatalf("Expected: %s, Got: %s", target, ret)
	}
}
