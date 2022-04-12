package user

import (
	"github.com/w33h/Hexagonal-Architecture-Go/business/user"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) user.Repository {
	return &postgresRepository{
		db: db,
	}
}

//func (repo *postgresRepository) Save(user user.Users) (user.Users, error) {
//	err := repo.db.Create(&user)
//	if err != nil {
//		return user, err.Error
//	}
//
//	return user, nil
//}
//
//func (repo *postgresRepository) Update(user user.Users) (user.Users, error) {
//	user.UpdatedAt = time.Now()
//
//	err := repo.db.Save(&user)
//	if err != nil {
//		return user, err.Error
//	}
//
//	return user, nil
//}
//
//func (repo *postgresRepository) Delete(Id int) (err error) {
//	var user user.Users
//
//	err = repo.db.Where("Id = ?", Id).Delete(&user).Error
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func (repo *postgresRepository) FindById(Id int) (user *user.Users, err error) {
	err = repo.db.Where("Id = ?", Id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

//func (repo *postgresRepository) FindAll() (users []user.Users, err error) {
//	err = repo.db.Find(&users).Error
//	if err != nil {
//		return nil, err
//	}
//
//	return users, nil
//}
