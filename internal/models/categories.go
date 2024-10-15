package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" gorm:"unique;size:200"`
	Label      string `json:"label" gorm:"size:200"`
	Position   uint   `json:"position"`
	ParentID   uint   `json:"parentId" gorm:"index"`
}

type CategoryWithPost struct {
	gorm.Model `json:"-"`
	CategoryID uint `json:"categoryId" gorm:"index"`
	PostID     uint `json:"postId" gorm:"index"`
}

func GetCategories(db *gorm.DB, parentID *uint) ([]Category, error) {
	var categories []Category
	tx := db
	if parentID != nil {
		tx = tx.Where("parent_id", *parentID)
	}
	if err := tx.Order("position DESC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func CountCategories(db *gorm.DB, parentID *uint) (total int64) {
	tx := db
	if parentID != nil {
		tx = tx.Where("parent_id", *parentID)
	}
	tx.Model(&Category{}).Count(&total)
	return
}
