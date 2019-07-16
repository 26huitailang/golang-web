package main_test

import (
	"golang_web/models"
	"golang_web/views"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestThemesHandle(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/themes", nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &views.Handler{Store: &StubStore{
		themes: []models.Theme{
			{
				Name: "Hey Ha",
			},
		},
	}}

	// Assertions
	if assert.NoError(t, h.ThemesHandle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, "Hey Ha", rec.Body.String()) // todo: bad test
	}
}

type StubStore struct {
	themes []models.Theme
}

func (s *StubStore) GetThemes() []models.Theme {
	return s.themes
}
