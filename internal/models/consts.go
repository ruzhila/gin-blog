package models

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DefaultLimit = 20

const (
	PostStatusDraft     = "draft"
	PostStatusPublished = "published"
)
const (
	Key_SiteName        = "SiteName"
	Key_SiteLogo        = "SiteLogo"
	Key_SiteIcon        = "SiteIcon"
	Key_SiteUrl         = "SiteUrl"
	Key_SiteAdmin       = "SiteAdmin"
	Key_SiteAdminEmail  = "SiteAdminEmail"
	Key_SiteTheme       = "SiteTheme"
	Key_SiteLang        = "SiteLang"
	Key_SiteKeywords    = "SiteKeywords"
	Key_SiteDescription = "SiteDescription"
	Key_SiteCopyRight   = "SiteCopyRight"
	Key_SiteGA          = "SiteGA"
	Key_SiteICP         = "SiteICP"
	Key_AllowRegistion  = "AllowRegistion"
)

func MakeMigration(db *gorm.DB) error {
	return db.AutoMigrate(&Post{},
		&Tag{},
		&Category{},
		&Comment{},
		&User{},
		&Config{},
		&PostPageView{},
		&Media{},
	)
}

func OpenDatabase(driver, dsn string) (db *gorm.DB, err error) {
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

	err = MakeMigration(db)
	if err != nil {
		return nil, err
	}
	logrus.WithField("driver", driver).WithField("dsn", dsn).Info("database connected")
	return db, nil
}
