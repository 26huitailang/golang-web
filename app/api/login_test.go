package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/26huitailang/golang_web/library/response"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// Setup
	e := echo.New()
	// TODO 这里的db没有mock，能不能吧testify的suite嵌套用于mockdb
	userJSON := `{"username": "test", "password": "123123123123"}`
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")
	c.SetRequest(req)

	// Assertions
	if assert.NoError(t, Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp response.JsonResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.Equal(t, response.OK, resp.Code)
		assert.Equal(t, "ok", resp.Message)
	}
}
