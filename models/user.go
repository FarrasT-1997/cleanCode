package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}

type Book struct {
	gorm.Model
	Book    string `json:"book" form:"book"`
	Country string `json:"country" form:"country"`
	Author  string `json:"author" form:"author"`
}
