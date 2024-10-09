package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handlers struct {
	db *gorm.DB
}

func (h *Handlers) Register(db *gorm.DB, engine *gin.Engine) error {
	return nil
}
