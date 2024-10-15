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

func CountPosts(db *gorm.DB) (total int64) {
	db.Model(&Post{}).Count(&total)
	return
}

func (p *Post) AfterFind(tx *gorm.DB) (err error) {
	tx = tx.Model(&PostPageView{}).Where("post_id", p.ID)
	err = tx.Select("SUM(page_view)").Scan(&p.PageView).Error
	if err != nil {
		return
	}
	var userView int64
	err = tx.Group("track_id").Count(&userView).Error
	p.UserView = uint(userView)
	return
}

func GetPosts(db *gorm.DB, offset, limit int) (posts []Post, err error) {
	if limit == 0 {
		limit = DefaultLimit
	}

	tx := db.Offset(offset).Limit(limit).Order("updated_at DESC")
	tx = tx.Preload("Tags").Preload("Author").Preload("Comments")
	err = tx.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return
}
