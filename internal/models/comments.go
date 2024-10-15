package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model `json:"-"`
	PostID     uint   `json:"-"`
	AuthorID   uint   `json:"-" gorm:"index"`
	Guest      bool   `json:"guest" gorm:"default:false"`
	GuestName  string `json:"guestName" gorm:"size:200"`
	Content    string `json:"content"`
}
