package usercontroller

import (
	"foodx-server/domain"
	"foodx-server/repository/user"
)

type UserController struct {
	userRepo *user.Service
}

func NewUserController(userRepo *user.Service) *UserController {
	return &UserController{userRepo: userRepo}
}

func (uc *UserController) GetUsers()(*[]domain.User, error){
	users, err := uc.userRepo.GetUsers()

	if err != nil{
		return nil, err
	}

	return users, nil
}

func (uc *UserController) GetUser(userID any)(*domain.User, error){
	user, err := uc.userRepo.GetUser(userID)
	
	if err != nil{
		return nil, err
	}

	return user, nil
}

func (uc *UserController) CraeteUser(user domain.User)(*domain.User, error){
	createdUser, err := uc.userRepo.CraeteUser(user)
	
	if err != nil{
		return nil, err
	}

	return createdUser, nil

}

func (uc *UserController) DeleteUser(userID any)(error){

	err := uc.userRepo.DeleteUser(userID)
	
	if err != nil{
		return err
	}

	return nil

}

func (uc *UserController) UserLogin(email string, password string) (*domain.UserLoginResponse, error){

	response, err := uc.userRepo.UserLogin(email,password)
	
	if err != nil{
		return nil, err
	}

	return  response, nil

}