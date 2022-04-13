package service

import (
	validator "github.com/go-playground/validator/v10"
	domain "github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"github.com/w33h/Hexagonal-Architecture-Go/business/user/spec"
	repository "github.com/w33h/Hexagonal-Architecture-Go/repository/user"
)

type service struct {
	repository repository.UserRepository
	validator  *validator.Validate
}

func NewService(repository repository.UserRepository) UserService {
	return &service{
		repository: repository,
		validator:  validator.New(),
	}
}
func (s service) FindById(userId int) (*domain.Users, error) {
	var user domain.Users

	userDB, err := s.repository.FindById(userId)

	if err != nil {
		return nil, err
	}

	user.Id = userDB.Id
	user.Email = userDB.Email
	user.Password = userDB.Password

	return &user, nil
}

func (s service) Create(user spec.UpsertUserSpec) (domain.Users, error) {
	err := s.validator.Struct(user)

	if err != nil {
		return domain.Users{}, err
	}
	var userInput domain.Users
	userInput.Email = user.Email
	userInput.Password = user.Password

	userDB, err := s.repository.Save(userInput)

	if err != nil {
		return domain.Users{}, err
	}

	return userDB, nil
}
