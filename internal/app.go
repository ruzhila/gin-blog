package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ruzhila/gin-blog/internal/handlers"
	"github.com/ruzhila/gin-blog/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BlogApp struct {
	db       *gorm.DB
	handlers *handlers.Handlers
}

func ConnectDB(driver, dsn string) (db *gorm.DB, err error) {
	driver = strings.ToLower(driver)
	cfg := &gorm.Config{}
	switch driver {
	case "sqlite3", "", "sqlite":
		if dsn == "" {
			dsn = "file::memory:"
		}
		db, err = gorm.Open(sqlite.Open(dsn), cfg)
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), cfg)
	case "postgres", "postgresql", "pg":
		db, err = gorm.Open(postgres.Open(dsn), cfg)
	default:
		err = fmt.Errorf("unsupported driver %s", driver)
	}
	if err != nil {
		return nil, err
	}

	err = models.MakeMigration(db)
	if err != nil {
		return nil, err
	}
	logrus.WithField("driver", driver).WithField("dsn", dsn).Info("database connected")
	return db, nil
}

func HintThemePath(themePath string) (string, bool) {
	hintTheme := false
	for _, d := range []string{".", "..", "../.."} {
		d = filepath.Join(d, themePath)
		if _, err := os.Stat(d); err == nil {
			themePath = d
			hintTheme = true
			// update theme path
			models.GetEnvs().ThemePath = themePath
			break
		}
	}
	return themePath, hintTheme
}

func NewBlogApp(db *gorm.DB) *BlogApp {
	return &BlogApp{
		db:       db,
		handlers: handlers.NewHandlers(db),
	}
}

func (app *BlogApp) Prepare(engine *gin.Engine) error {
	return app.handlers.Register(engine)
}
