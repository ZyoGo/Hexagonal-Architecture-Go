package service

import (
	domain "github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"github.com/w33h/Hexagonal-Architecture-Go/business/user/spec"
)

type UserService interface {
	Create(user spec.UpsertUserSpec) (domain.Users, error)
	//Update(user Users) (user.UserResponse, error)
	//Delete(userId int) error
	FindById(userId int) (*domain.Users, error)
	//FindAll() (users []user.UserResponse, err error)
}
