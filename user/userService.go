package user

import "errors"

type UserServiceInterface interface {
	GetAll() ([]Users, error)
	FindById(Id int) (Users, error)
	CreateUser(users UserRequest) (Users, error)
	UpdateUser(Id int, users UserRequest) (Users, error)
	DeleteUser(Id int) (bool, error)
}

type userService struct {
	userRepository UserRepositoryInterface
}

func NewUserService(repositoryUser UserRepositoryInterface) *userService {
	return &userService{repositoryUser}
}

func (service *userService) GetAll() ([]Users, error) {
	return service.userRepository.GetAll()
}

func (repository *userService) FindById(Id int) (Users, error) {
	return repository.userRepository.FindById(Id)
}

func (service *userService) CreateUser(users UserRequest) (Users, error) {
	insUser := Users{
		Username: users.Username,
		Name:     users.Name,
		Password: users.Password,
		Email:    users.Email,
		Active:   users.Active,
	}
	return service.userRepository.CreateUser(insUser)
}

func (service *userService) UpdateUser(Id int, users UserRequest) (Users, error) {
	findUser, err := service.FindById(Id)
	if err != nil {
		panic(err)
	}

	updateUser := Users{
		Id:       findUser.Id,
		Username: users.Username,
		Name:     users.Name,
		Password: users.Password,
		Email:    users.Email,
		Active:   users.Active,
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

func isEmptyStruck(user Users) bool {
	return user == Users{}
}
