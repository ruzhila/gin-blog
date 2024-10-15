package models

import (
	"gorm.io/gorm"
)

const (
	PostStatusDraft     = "draft"
	PostStatusPublished = "published"
)

func MakeMigration(db *gorm.DB) error {
	return db.AutoMigrate(&Post{},
		&Tag{},
		&Category{},
		&Comment{},
		&User{},
	)
}
