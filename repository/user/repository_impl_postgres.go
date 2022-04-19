package repository

import (
	"github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) user.UserRepository {
	return &userPostgresRepository{
		db: db,
	}
}

func (repo *userPostgresRepository) Save(user user.Users) (*user.Users, error) {
	err := repo.db.Create(&user)

	if err != nil {
		return nil, err.Error
	}

	return &user, nil
}

func (repo *userPostgresRepository) Update(user user.Users) (*user.Users, error) {
	err := repo.db.Save(&user)

	if err != nil {
		return nil, err.Error
	}

	return &user, nil
}

func (repo *userPostgresRepository) Delete(Id int) (err error) {
	if err = repo.db.Delete(&user.Users{}, "Id = ?", Id).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userPostgresRepository) FindById(Id int) (*user.Users, error) {
	err := repo.db.Where("Id = ?", Id)

	if err != nil {
		return nil, err.Error
	}

	var user user.Users
	user.Id = Id

	return &user, nil
}

func (repo *userPostgresRepository) FindAll() ([]user.Users, error) {
	var users []user.Users

	err := repo.db.Find(&user.Users{})

	if err != nil {
		return nil, err.Error
	}

	return users, nil
}

func (repo *postgresRepository) Save(user domain.Users) (domain.Users, error) {
	err := repo.db.Create(&user).Error
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}
