package employee

import (
	"errors"
	"time"
)

type employeeServiceInterface interface {
	GetAll() ([]Employee, error)
	FindById(Id int) (Employee, error)
	CreateEmployee(employee EmployeeRequest) (Employee, error)
	UpdateEmployee(Id int, employee EmployeeRequest) (Employee, error)
	DeleteEmployee(Id int) (bool, error)
}

type employeeService struct {
	employeRepository EmployeeRepositoryInterafce
}

func NewEmployeeService(repository EmployeeRepositoryInterafce) *employeeService {
	return &employeeService{repository}
}

func (service *employeeService) GetAll() ([]Employee, error) {
	return service.employeRepository.GetAll()
}

func (service *employeeService) FindById(Id int) (Employee, error) {
	return service.employeRepository.FindById(Id)
}

func (service *employeeService) CreateEmployee(request EmployeeRequest) (Employee, error) {
	insertEmployee := Employee{
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

func (service *employeeService) UpdateEmployee(Id int, request EmployeeRequest) (Employee, error) {
	findEmployee, err := service.employeRepository.FindById(Id)
	if err != nil {
		return Employee{}, errors.New("Data Tidak Ditemukan")
	}

	updateEmployee := Employee{
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
