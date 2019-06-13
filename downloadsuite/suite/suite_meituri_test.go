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

func TestFindSuitePageMax(t *testing.T) {
	input := []string{
		`<center><div id="pages"><a class="a1" href="https://www.meituri.com/a/26718/">上一页</a> <span>1</span> <a href="https://www.meituri.com/a/26718/2.html">2</a> <a href="https://www.meituri.com/a/26718/3.html">3</a> <a href="https://www.meituri.com/a/26718/4.html">4</a> <a href="https://www.meituri.com/a/26718/5.html">5</a> <a href="https://www.meituri.com/a/26718/6.html">6</a> <a href="https://www.meituri.com/a/26718/7.html">7</a> <a href="https://www.meituri.com/a/26718/8.html">8</a> <a href="https://www.meituri.com/a/26718/9.html">9</a> <a href="https://www.meituri.com/a/26718/10.html">10</a> ..<a href="https://www.meituri.com/a/26718/12.html">12</a> <a class="a1" href="https://www.meituri.com/a/26718/2.html">下一页</a></div></center>
		`,
		"",
	}
	output := []int{12, 1}
	for i, content := range input {
		out := findSuitePageMax(content)
		if out != output[i] {
			t.Fatalf("Expected %d, Got %d", output[i], out)
		}
	}
}
