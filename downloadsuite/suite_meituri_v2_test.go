package downloadsuite_test

import (
	"github.com/stretchr/testify/assert"
	"golang_web/downloadsuite"
	"testing"
)

func TestMeituriSuite(t *testing.T) {
	t.Run("获取suite所有的URL", func(t *testing.T) {

	})

	t.Run("获取title", func(t *testing.T) {
		operator := downloadsuite.NewMeituriSuite("mtr_suite.html", "/tmp", &downloadsuite.StubParser{})
		assert.Equal(t, "黑丝亮皮连衣超短裙 [森萝财团] [BETA-038] 写真集", operator.Title)
	})

	t.Run("获取拍摄机构URL", func(t *testing.T) {
		operator := downloadsuite.NewMeituriSuite("mtr_suite.html", "/tmp", &downloadsuite.StubParser{})
		assert.Equal(t, "https://www.meituri.com/x/82/", operator.OrgURL)
	})

	t.Run("获取suite最大页码", func(t *testing.T) {
		operator := downloadsuite.NewMeituriSuite("mtr_suite.html", "/tmp", &downloadsuite.StubParser{})
		assert.Equal(t, 12, operator.PageMax)

		operator2 := downloadsuite.NewMeituriSuite("mtr_suite_without_page.html", "/tmp", &downloadsuite.StubParser{})
		assert.Equal(t, 1, operator2.PageMax)
	})
}
