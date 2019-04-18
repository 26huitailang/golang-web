package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/26huitailang/golang-web/routers"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	r := gin.Default()
	baseFolder := "./static"
	routers.SetupRouter(r, baseFolder)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
	Convey("测试/ping接口", t, func() {
		So(w.Body.String(), ShouldEqual, "pong")
	})
}
