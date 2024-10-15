package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model `json:"-"`
	Slug       string     `json:"slug" gorm:"unique;size:200"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	AuthorID   uint       `json:"-"`
	Author     User       `json:"author,omitempty"`
	Published  bool       `json:"published" gorm:"default:false;index"`
	Tags       []Tag      `json:"tags,omitempty"`
	Categories []Category `json:"categories,omitempty"`
	Comments   []Comment  `json:"comments,omitempty"`
}

type PostLog struct {
	gorm.Model `json:"-"`
	PostID     uint   `json:"postId" gorm:"index"`
	Body       string `json:"body"`
}

type Tag struct {
	gorm.Model `json:"-"`
	PostID     uint   `json:"-" gorm:"index"`
	Name       string `json:"name" gorm:"unique;size:200"`
	Label      string `json:"label" gorm:"size:200"`
}

type Category struct {
	gorm.Model `json:"-"`
	PostID     uint   `json:"postId" gorm:"index"`
	Name       string `json:"name" gorm:"unique;size:200"`
	Label      string `json:"label" gorm:"size:200"`
}
