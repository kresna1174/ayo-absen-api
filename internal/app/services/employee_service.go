package services

import (
	"api-ayo-absen/internal/app/models"
	"api-ayo-absen/internal/app/repositories"
	"api-ayo-absen/internal/app/request"
	"errors"
)

type EmployeeServiceInterface interface {
	GetAll() ([]models.EmployeeWithCompany, error)
	FindById(Id int) (models.EmployeeWithCompany, error)
	CreateEmployee(employee request.EmployeeRequest) (models.EmployeeWithCompany, error)
	UpdateEmployee(Id int, employee request.EmployeeUpdateRequest) (models.EmployeeWithCompany, error)
	DeleteEmployee(Id int) (bool, error)
}

type employeeService struct {
	employeRepository repositories.EmployeeRepositoryInterafce
}

func NewEmployeeService(repository repositories.EmployeeRepositoryInterafce) *employeeService {
	return &employeeService{repository}
}

func (service *employeeService) GetAll() ([]models.EmployeeWithCompany, error) {
	return service.employeRepository.GetAll()
}

func (service *employeeService) FindById(Id int) (models.EmployeeWithCompany, error) {
	return service.employeRepository.FindById(Id)
}

func (service *employeeService) CreateEmployee(request request.EmployeeRequest) (models.EmployeeWithCompany, error) {
	insertEmployee := models.Employee{
		CompanyId: request.CompanyId,
		Name:      request.Name,
		Start:     request.Start,
		End:       request.End,
		Active:    request.Active,
		CreatedAt: request.CreatedAt,
		CreatedBy: request.CreatedBy,
		UpdatedAt: request.CreatedAt,
	}

	return service.employeRepository.CreateEmployee(insertEmployee)
}

func (service *employeeService) UpdateEmployee(Id int, request request.EmployeeUpdateRequest) (models.EmployeeWithCompany, error) {
	findEmployee, err := service.employeRepository.FindById(Id)
	if err != nil {
		return models.EmployeeWithCompany{}, errors.New("Data Tidak Ditemukan")
	}

	updateEmployee := models.Employee{
		Id:        findEmployee.Id,
		CompanyId: request.CompanyId,
		Name:      request.Name,
		Start:     request.Start,
		End:       request.End,
		Active:    request.Active,
		UpdatedAt: request.UpdatedAt,
		UpdatedBy: request.UpdatedBy,
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
