package services

import (
	"api-ayo-absen/internal/app/models"
	"api-ayo-absen/internal/app/repositories"
	"api-ayo-absen/internal/app/request"
)

type AuthServiceInterface interface {
	SignUp(user request.UserRequest) (models.Users, error)
	Login(Id int) (models.Users, error)
	FindUsername(username string) (models.Users, error)
}

type authService struct {
	authRepository repositories.AuthRepositoryInterface
}

func NewAuthService(authRepository repositories.AuthRepositoryInterface) *authService {
	return &authService{authRepository}
}

func (s *authService) SignUp(user request.UserRequest) (models.Users, error) {
	data := models.Users{
		Username: user.Username,
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
		Active:   1,
	}
	return s.authRepository.SignUp(data)
}

func (s *authService) Login(Id int) (models.Users, error) {
	return s.authRepository.Login(Id)
}

func (s *authService) FindUsername(username string) (models.Users, error) {
	return s.authRepository.FindUsername(username)
}
