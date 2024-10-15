package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) handleIndexPage(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
