package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/ruzhila/gin-blog/internal/handlers"
	"gorm.io/gorm"
)

type BlogApp struct {
	db       *gorm.DB
	handlers handlers.Handlers
}

func NewBlogApp(db *gorm.DB) *BlogApp {
	return &BlogApp{
		db:       db,
		handlers: handlers.Handlers{},
	}
}

func (app *BlogApp) Prepare(engine *gin.Engine) error {
	return app.handlers.Register(app.db, engine)
}
