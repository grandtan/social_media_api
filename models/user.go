package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"unique;type:varchar(100);not null"`
	Posts []Post
}
