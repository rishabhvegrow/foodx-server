package domain

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name     string `gorm:"not null"`
    Email    string `gorm:"not null"`
    Address  string
    Contact  string `gorm:"not null"`
    Password string `gorm:"not null"`
}

type UserLoginResponse struct {
    Token string `json:"token"`
    Email string `json:"email"`
    Name  string `json:"name"`
}