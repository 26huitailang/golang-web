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
		suite := downloadsuite.NewMeituriSuite("https://www.meituri.com/a/26718/")
		assert.Equal(t, suite.Title, "黑丝亮皮连衣超短裙 [森萝财团] [BETA-038] 写真集")
	})
}
