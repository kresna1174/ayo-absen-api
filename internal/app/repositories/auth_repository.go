package repositories

import (
	"api-ayo-absen/internal/app/models"

	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	SignUp(user models.Users) (models.Users, error)
	Login(Id int) (models.Users, error)
	FindUsername(username string) (models.Users, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) SignUp(user models.Users) (models.Users, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *authRepository) Login(Id int) (models.Users, error) {
	var user models.Users
	err := r.db.Find(&user, Id).Error

	return user, err
}

func (r *authRepository) FindUsername(username string) (models.Users, error) {
	var user models.Users
	err := r.db.First(&user, "username = ?", username).Error

	return user, err
}
