package repository

import "github.com/w33h/Hexagonal-Architecture-Go/business/user"

// Ingoing Port
type UserRepository interface {
	//Save(user Users) (Users, error)
	//Update(user Users) (Users, error)
	//Delete(Id int) (err error)
	FindById(Id int) (user *domain.Users, err error)
	//FindAll() (users []Users, err error)
}