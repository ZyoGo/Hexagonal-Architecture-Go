package user

import "github.com/w33h/Hexagonal-Architecture-Go/business/user/spec"

type UserService interface {
	Create(userSpec spec.UpsertUserSpec) (Users, error)
	Update(userSpec spec.UpsertUserSpec) (Users, error)
	Delete(id int) error
	FindById(id int) (Users, error)
	FindAll() (users []Users, err error)
}
