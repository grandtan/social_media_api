package models

import "gorm.io/gorm"

// Post struct defines the post model
type Post struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Content string `gorm:"type:text"`
}
