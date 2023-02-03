package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Autor    string
	Name      string
	PageCount int
}
