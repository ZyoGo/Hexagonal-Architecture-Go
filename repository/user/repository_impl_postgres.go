package user

import (
	domain "github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) UserRepository {
	return &userPostgresRepository{
		db: db,
	}
}

func (repo *userPostgresRepository) Save(user domain.Users) (*domain.Users, error) {
	err := repo.db.Create(&user)

	if err != nil {
		return nil, err.Error
	}

	return &user, nil
}

func (repo *userPostgresRepository) Update(user domain.Users) (*domain.Users, error) {
	err := repo.db.Save(&user)

	if err != nil {
		return nil, err.Error
	}

	return &user, nil
}

func (repo *userPostgresRepository) Delete(Id int) (err error) {
	if err = repo.db.Delete(&domain.Users{}, "Id = ?", Id).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userPostgresRepository) FindById(Id int) (*domain.Users, error) {
	err := repo.db.Where("Id = ?", Id)

	if err != nil {
		return nil, err.Error
	}

	var user domain.Users
	user.Id = Id

	return &user, nil
}

func (repo *userPostgresRepository) FindAll() ([]domain.Users, error) {
	var users []domain.Users

	err := repo.db.Find(&domain.Users{})

	if err != nil {
		return nil, err.Error
	}

	return users, nil
}
