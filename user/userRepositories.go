package user

import (
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetAll() ([]Users, error)
	FindById(Id int) (Users, error)
	CreateUser(users Users) (Users, error)
	UpdateUser(user Users) (Users, error)
	DeleteUser(user Users) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (repository *userRepository) GetAll() ([]Users, error) {
	var users []Users

	err := repository.db.Find(&users).Error

	return users, err
}

func (repository *userRepository) FindById(Id int) (Users, error) {
	var user Users

	err := repository.db.Find(&user, Id).Error

	return user, err
}

func (repository *userRepository) CreateUser(users Users) (Users, error) {
	err := repository.db.Create(&users).Error

	return users, err
}

func (repository *userRepository) UpdateUser(user Users) (Users, error) {
	err := repository.db.Save(&user).Error

	return user, err
}

func (repository *userRepository) DeleteUser(user Users) (bool, error) {
	err := repository.db.Delete(user).Error

	status := true
	if err != nil {
		status = false
	}

	return status, err
}
