package services

import (
	"api-ayo-absen/internal/app/models"
	"api-ayo-absen/internal/app/repositories"
	"api-ayo-absen/internal/app/request"
	"errors"
	"time"
)

type EmployeeServiceInterface interface {
	GetAll() ([]models.Employee, error)
	FindById(Id int) (models.Employee, error)
	CreateEmployee(employee request.EmployeeRequest) (models.Employee, error)
	UpdateEmployee(Id int, employee request.EmployeeRequest) (models.Employee, error)
	DeleteEmployee(Id int) (bool, error)
}

type employeeService struct {
	employeRepository repositories.EmployeeRepositoryInterafce
}

func NewEmployeeService(repository repositories.EmployeeRepositoryInterafce) *employeeService {
	return &employeeService{repository}
}

func (service *employeeService) GetAll() ([]models.Employee, error) {
	return service.employeRepository.GetAll()
}

func (service *employeeService) FindById(Id int) (models.Employee, error) {
	return service.employeRepository.FindById(Id)
}

func (service *employeeService) CreateEmployee(request request.EmployeeRequest) (models.Employee, error) {
	insertEmployee := models.Employee{
		CompanyId: request.CompanyId,
		Name:      request.Name,
		Start:     time.Now(),
		End:       time.Now(),
		Active:    request.Active,
		CreatedAt: time.Now(),
		CreatedBy: "system",
		UpdatedAt: time.Now(),
		UpdatedBy: "system",
	}

	return service.employeRepository.CreateEmployee(insertEmployee)
}

func (service *employeeService) UpdateEmployee(Id int, request request.EmployeeRequest) (models.Employee, error) {
	findEmployee, err := service.employeRepository.FindById(Id)
	if err != nil {
		return models.Employee{}, errors.New("Data Tidak Ditemukan")
	}

	updateEmployee := models.Employee{
		Id:        findEmployee.Id,
		CompanyId: request.CompanyId,
		Name:      request.Name,
		Start:     time.Now(),
		End:       time.Now(),
		Active:    request.Active,
		CreatedAt: findEmployee.CreatedAt,
		CreatedBy: findEmployee.CreatedBy,
		UpdatedAt: time.Now(),
		UpdatedBy: "system",
	}

	return service.employeRepository.UpdateEmployee(updateEmployee)
}

func (service *employeeService) DeleteEmployee(Id int) (bool, error) {
	findEmployee, err := service.employeRepository.FindById(Id)
	if err != nil {
		return false, errors.New("Data Tidak Ditemukan")
	}

	return service.employeRepository.DeleteEmployee(findEmployee)
}
