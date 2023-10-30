package usercontroller

import (
	"foodx-server/domain"
)

type Reader interface {
	GetUsers()(*[]domain.User, error)
	GetUser(userID any)(*domain.User, error)
	UserLogin(email string, password string) (*domain.UserLoginResponse, error)
}

type Writer interface {
	CraeteUser(user domain.User)(*domain.User, error)
	DeleteUser(userID any)(error)
}
type Controller interface {
	Reader
	Writer
}
