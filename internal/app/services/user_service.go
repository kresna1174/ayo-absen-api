package services

import (
	"api-ayo-absen/internal/app/models"
	"api-ayo-absen/internal/app/repositories"
	"api-ayo-absen/internal/app/request"
	"errors"
)

type UserServiceInterface interface {
	GetAll() ([]models.Users, error)
	FindById(Id int) (models.Users, error)
	CreateUser(users request.UserRequest) (models.Users, error)
	UpdateUser(Id int, users request.UserUpdateRequest) (models.Users, error)
	DeleteUser(Id int) (bool, error)
}

type userService struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserService(repositoryUser repositories.UserRepositoryInterface) *userService {
	return &userService{repositoryUser}
}

func (service *userService) GetAll() ([]models.Users, error) {
	return service.userRepository.GetAll()
}

func (repository *userService) FindById(Id int) (models.Users, error) {
	return repository.userRepository.FindById(Id)
}

func (service *userService) CreateUser(users request.UserRequest) (models.Users, error) {
	insUser := models.Users{
		Username:  users.Username,
		Name:      users.Name,
		Password:  users.Password,
		Email:     users.Email,
		Active:    users.Active,
		CreatedAt: users.CreatedAt,
		CreatedBy: users.CreatedBy,
	}
	return service.userRepository.CreateUser(insUser)
}

func (service *userService) UpdateUser(Id int, users request.UserUpdateRequest) (models.Users, error) {
	findUser, err := service.FindById(Id)
	if err != nil {
		panic(err)
	}

	updateUser := models.Users{
		Id:        findUser.Id,
		Username:  users.Username,
		Name:      users.Name,
		Email:     users.Email,
		Active:    users.Active,
		UpdatedAt: users.UpdatedAt,
		UpdatedBy: users.UpdatedBy,
	}

	if users.Password != "" {
		updateUser.Password = users.Password
	}

	return service.userRepository.UpdateUser(updateUser)
}

func (service *userService) DeleteUser(Id int) (bool, error) {
	findUser, err := service.FindById(Id)
	if err != nil {
		panic(err)
	}
	if isEmptyStruck(findUser) {
		return false, errors.New("data tidak ditemukan")
	}

	return service.userRepository.DeleteUser(findUser)
}

func isEmptyStruck(user models.Users) bool {
	return user == models.Users{}
}
