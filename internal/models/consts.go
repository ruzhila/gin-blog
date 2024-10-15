package models

import (
	"gorm.io/gorm"
)

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
)

func MakeMigration(db *gorm.DB) error {
	return db.AutoMigrate(&Post{},
		&Tag{},
		&Category{},
		&Comment{},
		&User{},
		&Config{},
	)
}
