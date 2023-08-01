package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

type UserInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
