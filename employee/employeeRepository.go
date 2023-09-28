package employee

import "gorm.io/gorm"

type EmployeeRepositoryInterafce interface {
	GetAll() ([]Employee, error)
	FindById(Id int) (Employee, error)
	CreateEmployee(employee Employee) (Employee, error)
	UpdateEmployee(employee Employee) (Employee, error)
	DeleteEmployee(employee Employee) (bool, error)
}

type employeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employeRepository {
	return &employeRepository{db}
}

func (repository *employeRepository) GetAll() ([]Employee, error) {
	var employee []Employee

	err := repository.db.Find(&employee).Error

	return employee, err
}

func (repository *employeRepository) FindById(Id int) (Employee, error) {
	var employee Employee

	err := repository.db.Find(&employee, Id).Error

	return employee, err
}

func (repository *employeRepository) CreateEmployee(employee Employee) (Employee, error) {
	err := repository.db.Create(&employee).Error

	return employee, err
}

func (repository *employeRepository) UpdateEmployee(employee Employee) (Employee, error) {
	err := repository.db.Save(&employee).Error

	return employee, err
}

func (repository *employeRepository) DeleteEmployee(employee Employee) (bool, error) {
	status := true
	err := repository.db.Delete(&employee).Error
	if err != nil {
		status = false
	}

	return status, err
}
