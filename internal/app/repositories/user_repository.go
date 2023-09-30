package repositories

import (
	"api-ayo-absen/internal/app/models"

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
	err := repository.db.Save(&user).Error

	return user, err
}

func (repository *userRepository) DeleteUser(user models.Users) (bool, error) {
	err := repository.db.Delete(user).Error

	status := true
	if err != nil {
		status = false
	}

	return status, err
}
