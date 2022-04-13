package service

import helper "github.com/w33h/Hexagonal-Architecture-Go/helpers/user"

type UserService interface {
	//Create(user Users) (user.U, error)
	//Update(user Users) (user.UserResponse, error)
	//Delete(userId int) error
	FindById(userId int) (helper.UserResponse, error)
	//FindAll() (users []user.UserResponse, err error)
}
