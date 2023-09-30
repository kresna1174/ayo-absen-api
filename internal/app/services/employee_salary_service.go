package services

import (
	"api-ayo-absen/internal/app/models"
	"api-ayo-absen/internal/app/repositories"
	"api-ayo-absen/internal/app/request"
	"errors"
	"time"
)

type EmployeeSalaryServiceInterface interface {
	GetAll() ([]models.EmployeeSalaries, error)
	FindById(Id int) (models.EmployeeSalaries, error)
	Create(employeeSalary request.EmployeeSalaryRequest) (models.EmployeeSalaries, error)
	Update(Id int, employeeSalary request.EmployeeSalaryRequest) (models.EmployeeSalaries, error)
	Delete(Id int) (bool, error)
}

type employeeSalaryService struct {
	employeeRepositoryInterface repositories.EmployeeSalaryRepositoryInterface
}

func NewEmployeeSalaryService(employeeSalaryRepositoryInterface repositories.EmployeeSalaryRepositoryInterface) *employeeSalaryService {
	return &employeeSalaryService{employeeSalaryRepositoryInterface}
}

func (service *employeeSalaryService) GetAll() ([]models.EmployeeSalaries, error) {
	return service.employeeRepositoryInterface.GetAll()
}

func (service *employeeSalaryService) FindById(Id int) (models.EmployeeSalaries, error) {
	return service.employeeRepositoryInterface.FindById(Id)
}

func (service *employeeSalaryService) Create(employeeSalary request.EmployeeSalaryRequest) (models.EmployeeSalaries, error) {
	employeeSalaryEntity := models.EmployeeSalaries{
		CompanyId:  employeeSalary.CompanyId,
		EmployeeId: employeeSalary.EmployeeId,
		Salary:     employeeSalary.Salary,
		PayPeriod:  employeeSalary.PayPeriod,
		Created_at: time.Now(),
		Created_by: "system",
		Updated_at: time.Now(),
		Updated_by: "system",
	}
	return service.employeeRepositoryInterface.Create(employeeSalaryEntity)
}

func (service *employeeSalaryService) Update(Id int, employee request.EmployeeSalaryRequest) (models.EmployeeSalaries, error) {
	findEmployeeSalary, err := service.employeeRepositoryInterface.FindById(Id)
	if err != nil {
		return models.EmployeeSalaries{}, errors.New("Data tidak ditemukan")
	}

	employeeSalaryEntity := models.EmployeeSalaries{
		Id:         findEmployeeSalary.Id,
		CompanyId:  employee.CompanyId,
		EmployeeId: employee.EmployeeId,
		Salary:     employee.Salary,
		PayPeriod:  employee.PayPeriod,
		Created_at: findEmployeeSalary.Created_at,
		Created_by: findEmployeeSalary.Created_by,
		Updated_at: time.Now(),
		Updated_by: "system",
	}
	return service.employeeRepositoryInterface.Update(employeeSalaryEntity)
}

func (service *employeeSalaryService) Delete(Id int) (bool, error) {
	findEmployeeSalary, err := service.employeeRepositoryInterface.FindById(Id)
	if err != nil {
		return false, errors.New("Data tidak ditemukan")
	}
	return service.employeeRepositoryInterface.Delete(findEmployeeSalary)
}
