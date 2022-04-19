package service

import (
	"github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"github.com/w33h/Hexagonal-Architecture-Go/business/user/spec"

	validator "github.com/go-playground/validator/v10"
)

type service struct {
	repo     user.UserRepository
	validate *validator.Validate
}

func NewServiceUser(repository user.UserRepository) user.UserService {
	return &service{
		repo:     repository,
		validate: validator.New(),
	}
}

func (s *service) Create(userSpec spec.UpsertUserSpec) (user.Users, error) {
	err := s.validate.Struct(&userSpec)
	if err != nil {
		return user.Users{}, err
	}

	var user user.Users
	user.Username = userSpec.Username
	user.Email = userSpec.Email
	user.Password = userSpec.Password

	result, err := s.repo.Save(user)

	if err != nil {
		return *result, err
	}

	return *result, nil
}

func (s *service) Update(userSpec spec.UpsertUserSpec) (user.Users, error) {
	err := s.validate.Struct(&userSpec)
	if err != nil {
		return user.Users{}, err
	}

	var user user.Users
	user.Username = userSpec.Username
	user.Email = userSpec.Email
	user.Password = userSpec.Password

	result, err := s.repo.Update(user)
	if err != nil {
		return *result, err
	}

	return *result, nil
}

func (s *service) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) FindById(id int) (user.Users, error) {
	result, err := s.repo.FindById(id)
	if err != nil {
		return *result, err
	}

	return *result, nil
}

func (s *service) FindAll() (users []user.Users, err error) {
	users, err = s.repo.FindAll()

	if err != nil {
		return users, err
	}

	return users, nil
}
