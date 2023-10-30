package user

import (
	"foodx-server/domain"
)

// Reader Interfce
type Reader interface {
	GetUsers()(*[]domain.User, error)
	GetUser(userID any)(*domain.User, error)
	UserLogin(email string, password string) (*domain.UserLoginResponse, error)
}

type Writer interface {
	CraeteUser(user domain.User)(*domain.User, error)
	DeleteUser(userID any)(error)
}

// Filter interface
type Filter interface {
}

// Repository interface
type Repository interface {
	Reader
	Writer
	Filter
}