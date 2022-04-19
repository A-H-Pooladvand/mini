package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Title string
	Body  string
}
