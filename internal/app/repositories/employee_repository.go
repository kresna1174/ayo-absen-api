package repositories

import (
	"api-ayo-absen/internal/app/models"
	"errors"

	"gorm.io/gorm"
)

type EmployeeRepositoryInterafce interface {
	GetAll() ([]models.EmployeeWithCompany, error)
	FindById(Id int) (models.EmployeeWithCompany, error)
	CreateEmployee(employee models.Employee) (models.Employee, error)
	UpdateEmployee(employee models.Employee) (models.Employee, error)
	DeleteEmployee(employee models.EmployeeWithCompany) (bool, error)
}

type employeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employeRepository {
	return &employeRepository{db}
}

func (repository *employeRepository) GetAll() ([]models.EmployeeWithCompany, error) {
	var employee models.Employee
	return employee.ViewCompany(repository.db)
}

func (repository *employeRepository) FindById(Id int) (models.EmployeeWithCompany, error) {
	var employee models.Employee
	return employee.ViewCompanySinggle(repository.db, Id)
	// var employee models.Employee

	// err := repository.db.Find(&employee, Id).Error

	// return employee, err
}

func (repository *employeRepository) CreateEmployee(employee models.Employee) (models.Employee, error) {
	err := repository.db.Create(&employee).Error

	return employee, err
}

func (repository *employeRepository) UpdateEmployee(employee models.Employee) (models.Employee, error) {
	err := repository.db.Model(&employee).Updates(&employee).Error
	var employeeModel models.Employee
	er := repository.db.Find(&employeeModel, employee.Id).Error
	if er != nil {
		return models.Employee{}, errors.New("tidak dapat mengembalikan nilai")
	}
	return employeeModel, err
}

func (repository *employeRepository) DeleteEmployee(employee models.EmployeeWithCompany) (bool, error) {
	status := true
	err := repository.db.Delete(&employee).Error
	if err != nil {
		status = false
	}

	return status, err
}
