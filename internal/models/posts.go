package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model `json:"-"`
	Slug       string    `json:"slug" gorm:"unique;size:200"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   uint      `json:"-"`
	Author     User      `json:"author,omitempty"`
	Published  bool      `json:"published" gorm:"default:false;index"`
	Tags       []Tag     `json:"tags,omitempty"`
	Comments   []Comment `json:"comments,omitempty"`
	PageView   uint      `json:"pageView" gorm:"-"`
	UserView   uint      `json:"userView" gorm:"-"`
}

type PostLog struct {
	gorm.Model `json:"-"`
	PostID     uint   `json:"postId" gorm:"index"`
	Body       string `json:"body"`
}

type PostPageView struct {
	ID       uint      `json:"id"`
	PostID   uint      `json:"postId" gorm:"unique_index:idx_post_id_when_track"`
	When     time.Time `json:"when" gorm:"unique_index:idx_post_id_when_track"`
	TrackID  string    `json:"trackId" gorm:"size:64;unique_index:idx_post_id_when_track"`
	IP       string    `json:"ip" gorm:"size:64"`
	PageView uint      `json:"pageView"`
}

type Tag struct {
	gorm.Model `json:"-"`
	PostID     uint   `json:"-" gorm:"index"`
	Name       string `json:"name" gorm:"unique;size:200"`
	Label      string `json:"label" gorm:"size:200"`
}

type Category struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" gorm:"unique;size:200"`
	Label      string `json:"label" gorm:"size:200"`
	Sort       int    `json:"sort"`
	ParentID   uint   `json:"parentId" gorm:"index"`
}

type CategoryWithPost struct {
	gorm.Model `json:"-"`
	CategoryID uint `json:"categoryId" gorm:"index"`
	PostID     uint `json:"postId" gorm:"index"`
}
