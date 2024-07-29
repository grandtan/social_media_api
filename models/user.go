package models

import "gorm.io/gorm"

// User struct defines the user model
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"unique;type:varchar(100);not null"`
	Posts []Post
}
