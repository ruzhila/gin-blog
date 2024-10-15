package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/ruzhila/gin-blog/internal/handlers"
	"github.com/ruzhila/gin-blog/internal/models"
	"github.com/ruzhila/gin-blog/internal/themes"
	"gorm.io/gorm"
)

type BlogApp struct {
	db       *gorm.DB
	handlers *handlers.Handlers
}

func NewBlogApp(db *gorm.DB) *BlogApp {
	return &BlogApp{
		db:       db,
		handlers: handlers.NewHandlers(db),
	}
}

func (app *BlogApp) Prepare(engine *gin.Engine) error {
	models.CheckDefaultConfigValues(app.db)

	if p, ok := models.GetConfigValue(app.db, models.Key_SiteTheme); ok {
		app.handlers.Theme = p
	}

	r, err := themes.NewRender()
	if err != nil {
		return err
	}
	engine.HTMLRender = r
	return app.handlers.Register(engine)
}
