package models

import "gorm.io/gorm"

type User struct {
	gorm.Model  `json:"-"`
	Email       string `json:"-" gorm:"unique;size:200"`
	Password    string `json:"-" gorm:"size:128"`
	DisplayName string `json:"displayName" gorm:"size:2000"`
	Avatar      string `json:"avatar" gorm:"size:2000"`
	Location    string `json:"location" gorm:"size:200"`
	IsStaff     bool   `json:"-" gorm:"default:false"`
	IsSuper     bool   `json:"-" gorm:"default:false"`
	Disabled    bool   `json:"-" gorm:"default:false"`
}
