package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title   string `gorm:"unique"`
	Excerpt string
	Body    string
	Image   string
}
