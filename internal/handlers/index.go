package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"github.com/ruzhila/gin-blog/internal/models"
)

func (h *Handlers) themeFile(name string) string {
	return filepath.Join("themes", h.Theme, name)
}

func (h *Handlers) CTX(c gin.H) pongo2.Context {
	m := pongo2.Context{}
	for k, v := range models.GetConfigValues(h.db) {
		m[k] = v
	}
	for k, v := range c {
		m[k] = v
	}
	return WithFuncs(h.db, m)
}

func (h *Handlers) handleIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, h.themeFile("index.tpl"), h.CTX(nil))
}

func (h *Handlers) handleSitemap(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
