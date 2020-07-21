package views

import (
	"golang_web/config"
	"golang_web/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestThemesHandle(t *testing.T) {
	// Setup
	e := echo.New()
	e.Renderer = &config.Template{}
	req := httptest.NewRequest(http.MethodGet, "/themes", nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{Store: &StubStore{
		themes: []models.Theme{
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
	themes []models.Theme
}

func (s *StubStore) GetThemes() []models.Theme {
	return s.themes
}
