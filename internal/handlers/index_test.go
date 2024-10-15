package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ruzhila/gin-blog/internal/models"
	"github.com/ruzhila/gin-blog/internal/themes"
	"github.com/stretchr/testify/assert"
)

func createTestHandlers(w http.ResponseWriter) (*Handlers, *gin.Context) {
	db, _ := models.OpenDatabase("", "")
	models.CheckDefaultConfigValues(db)

	h := NewHandlers(db)
	c, r := gin.CreateTestContext(w)
	r.HTMLRender, _ = themes.NewRender()
	return h, c
}
func TestIndexPage(t *testing.T) {
	w := httptest.NewRecorder()
	h, c := createTestHandlers(w)
	h.handleIndexPage(c)
	assert.Equal(t, http.StatusOK, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, "<!DOCTYPE html>")
	assert.Contains(t, body, "<title>Gin blog</title>")
}
