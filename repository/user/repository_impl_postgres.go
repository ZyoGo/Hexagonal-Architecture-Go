package repository

import (
	domain "github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) UserRepository {
	return &postgresRepository{
		db: db,
	}
}

func (repo *postgresRepository) FindById(Id int) (user *domain.Users, err error) {
	err = repo.db.Where("Id = ?", Id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *postgresRepository) Save(user domain.Users) (domain.Users, error) {
	err := repo.db.Create(&user).Error
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}
