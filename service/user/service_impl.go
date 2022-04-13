package service

import (
	domain "github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"github.com/w33h/Hexagonal-Architecture-Go/helpers/user"
	repository "github.com/w33h/Hexagonal-Architecture-Go/repository/user"
)

type service struct {
	repository repository.UserRepository
}

func NewService(repository repository.UserRepository) UserService {
	return &service{
		repository: repository,
	}
}
func (s service) FindById(userId int) (helper.UserResponse, error) {
	var user domain.Users

	userDB, err := s.repository.FindById(userId)

	if err != nil {
		return helper.UserResponse{}, err
	}

	user.Id = userDB.Id
	user.Email = userDB.Email
	user.Password = userDB.Password

	return helper.ToUserResponse(*userDB), nil
}

