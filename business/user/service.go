package user

import "github.com/w33h/Hexagonal-Architecture-Go/helpers/helper"

type Service interface {
	//Create(user Users) (helper.UserResponse, error)
	//Update(user Users) (helper.UserResponse, error)
	//Delete(userId int) error
	FindById(userId int) (helper.UserResponse, error)
	//FindAll() (users []helper.UserResponse, err error)
}
