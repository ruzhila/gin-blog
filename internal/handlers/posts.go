package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) handlePost(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
func (h *Handlers) handleComment(c *gin.Context) {
}
