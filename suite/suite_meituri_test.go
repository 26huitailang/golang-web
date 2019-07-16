package suite_test

import (
	"golang_web/suite"
	"io/ioutil"
	"testing"
)

func TestParseOrgURL(t *testing.T) {
	target := "https://www.meituri.com/x/82/"
	data, err := ioutil.ReadFile("content.html")
	suite.CheckError(err)
	// todo: 不要测试内部函数
	ret := suite.ParseOrgURL(string(data))
	if ret != target {
		t.Fatalf("Expected: %s, Got: %s", target, ret)
	}

	// 多个标签，获取第一个
	input := `<p>拍摄机构：
		<a href="https://www.meituri.com/x/12/" target="_blank">异思趣向</a>
					<a href="https://www.meituri.com/x/14/" target="_blank">丝享家</a>
	</p>`
	output := "https://www.meituri.com/x/12/"
	ret2 := suite.ParseOrgURL(input)
	if ret2 != output {
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
		// todo: 不要测试内部函数
		out := suite.FindSuitePageMax(content)
		if out != output[i] {
			t.Fatalf("Expected %d, Got %d", output[i], out)
		}
	}
}
