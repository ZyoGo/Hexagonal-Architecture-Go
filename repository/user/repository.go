package user

import "github.com/w33h/Hexagonal-Architecture-Go/business/user"

// Ingoing Port
type UserRepository interface {
	Save(user user.Users) (*user.Users, error)
	Update(user user.Users) (*user.Users, error)
	Delete(Id int) (err error)
	FindById(Id int) (*user.Users, error)
	FindAll() ([]user.Users, error)
}
