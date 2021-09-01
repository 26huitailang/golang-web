package views

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/config"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestThemesHandle(t *testing.T) {
	// Setup
	e := echo.New()
	e.Renderer = &config.Template{}
	config.ReloadTemplates()
	req := httptest.NewRequest(http.MethodGet, "/themes", nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{IThemeStore: &StubStore{
		themes: []model.Theme{
			{
				Name: "Hey Ha",
			},
		},
	}}

	// Assertions
	if assert.NoError(t, h.ThemesHandle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Hey Ha") // todo: bad test
	}
}

type StubStore struct {
	themes []model.Theme
}

func (s *StubStore) GetThemes() []model.Theme {
	return s.themes
}
