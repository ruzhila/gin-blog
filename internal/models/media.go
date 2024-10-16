package models

import "gorm.io/gorm"

type Media struct {
	gorm.Model `json:"-"`
	AuthorID   uint   `json:"-"`
	Author     User   `json:"author,omitempty"`
	Filename   string `json:"filename" gorm:"size:200"`
	Ext        string `json:"ext" gorm:"size:200"`
	Size       int64  `json:"size"`
	StorePath  string `json:"storePath" gorm:"size:200"`
}
