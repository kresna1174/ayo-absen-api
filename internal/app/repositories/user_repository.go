package repositories

import (
	"api-ayo-absen/internal/app/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetAll() ([]models.Users, error)
	FindById(Id int) (models.Users, error)
	CreateUser(users models.Users) (models.Users, error)
	UpdateUser(user models.Users) (models.Users, error)
	DeleteUser(user models.Users) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (repository *userRepository) GetAll() ([]models.Users, error) {
	var users []models.Users

	err := repository.db.Find(&users).Error

	return users, err
}

func (repository *userRepository) FindById(Id int) (models.Users, error) {
	var user models.Users

	err := repository.db.Find(&user, Id).Error

	return user, err
}

func (repository *userRepository) CreateUser(users models.Users) (models.Users, error) {
	err := repository.db.Create(&users).Error

	return users, err
}

func (repository *userRepository) UpdateUser(user models.Users) (models.Users, error) {
	err := repository.db.Model(&user).Updates(&user).Error

	var userModel models.Users
	er := repository.db.Find(&userModel, user.Id).Error
	if er != nil {
		return models.Users{}, errors.New("tidak dapat mengembalikan nilai")
	}
	return userModel, err
}

func (repository *userRepository) DeleteUser(user models.Users) (bool, error) {
	err := repository.db.Delete(user).Error

	status := true
	if err != nil {
		status = false
	}

	return status, err
}
