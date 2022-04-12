package user

import (
	domain "github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"github.com/w33h/Hexagonal-Architecture-Go/helpers/helper"
)

type service struct {
	repository domain.Repository
}

func NewService(repository domain.Repository) domain.Service {
	return &service{
		repository: repository,
	}
}

//func (s service) Create(user domain.Users) (helper.UserResponse, error) {
//	insertUser, err := s.repository.Save(user)
//
//	if err != nil {
//		return helper.UserResponse{}, err
//	}
//
//	return helper.UserResponse{Id: insertUser.Id, Email: insertUser.Email, Password: insertUser.Password}, nil
//}
//
//func (s service) Update(user domain.Users) (helper.UserResponse, error) {
//	panic("implement me")
//}
//
//func (s service) Delete(userId int) error {
//	panic("implement me")
//}

func (s service) FindById(userId int) (helper.UserResponse, error) {
	var user helper.UserResponse

	userDB, err := s.repository.FindById(userId)

	if err != nil {
		return helper.UserResponse{}, err
	}

	user.Id = userDB.Id
	user.Email = userDB.Email
	user.Password = userDB.Password

	return user, nil
}

//func (s service) FindAll() (users []helper.UserResponse, err error) {
//	panic("implement me")
//}

