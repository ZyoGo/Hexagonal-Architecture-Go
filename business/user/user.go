package user

import (
	"time"

	"github.com/w33h/Hexagonal-Architecture-Go/business/user/spec"
)

type Users struct {
	Id        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Ingoing Port
type UserRepository interface {
	Save(user Users) (*Users, error)
	Update(user Users) (*Users, error)
	Delete(Id int) (err error)
	FindById(Id int) (*Users, error)
	FindAll() ([]Users, error)
}

// Outgoing Port
type UserService interface {
	Create(userSpec spec.UpsertUserSpec) (Users, error)
	Update(userSpec spec.UpsertUserSpec) (Users, error)
	Delete(id int) error
	FindById(id int) (Users, error)
	FindAll() (users []Users, err error)
}
