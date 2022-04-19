package user

import (
	"github.com/w33h/Hexagonal-Architecture-Go/business/user/spec"
	repository "github.com/w33h/Hexagonal-Architecture-Go/repository/user"

	validator "github.com/go-playground/validator/v10"
)

type service struct {
	repo repository.UserRepository
	validate *validator.Validate

}

func NewService(repository repository.UserRepository) UserService {
	return &service{
		repo: repository,
		validate: validator.New(),
	}
}

func (s *service) Create(userSpec spec.UpsertUserSpec) (Users, error) {
	err := s.validate.Struct(&userSpec)
	if err != nil {
		return Users{}, err
	}

	var user Users
	user.Username = userSpec.Username
	user.Email = userSpec.Email
	user.Password = userSpec.Password

	result, err := s.repo.Save(user)

	if err != nil {
		return *result, err
	}

	return *result, nil
}

func (s *service) Update(userSpec spec.UpsertUserSpec) (Users, error) {
	err := s.validate.Struct(&userSpec)
	if err != nil {
		return Users{}, err
	}

	var user Users
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

func (s *service) FindById(id int) (Users, error) {
	result, err := s.repo.FindById(id)
	if err != nil {
		return *result, err
	}

	return *result, nil
}

func (s *service) FindAll() (users []Users, err error) {
	users, err = s.repo.FindAll()

	if err != nil {
		return users, err
	}

	return users, nil
}
