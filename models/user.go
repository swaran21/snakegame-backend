package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex;not null"`
	HighScore int    `gorm:"default:0"`
}
