package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Book    string `json:"book" form:"book"`
	Country string `json:"country" form:"country"`
	Author  string `json:"author" form:"author"`
}
