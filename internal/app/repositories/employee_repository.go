package repositories

import (
	"api-ayo-absen/internal/app/models"

	"gorm.io/gorm"
)

type EmployeeRepositoryInterafce interface {
	GetAll() ([]models.Employee, error)
	FindById(Id int) (models.Employee, error)
	CreateEmployee(employee models.Employee) (models.Employee, error)
	UpdateEmployee(employee models.Employee) (models.Employee, error)
	DeleteEmployee(employee models.Employee) (bool, error)
}

type employeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employeRepository {
	return &employeRepository{db}
}

func (repository *employeRepository) GetAll() ([]models.Employee, error) {
	var employee []models.Employee

	err := repository.db.Find(&employee).Error

	return employee, err
}

func (repository *employeRepository) FindById(Id int) (models.Employee, error) {
	var employee models.Employee

	err := repository.db.Find(&employee, Id).Error

	return employee, err
}

func (repository *employeRepository) CreateEmployee(employee models.Employee) (models.Employee, error) {
	err := repository.db.Create(&employee).Error

	return employee, err
}

func (repository *employeRepository) UpdateEmployee(employee models.Employee) (models.Employee, error) {
	err := repository.db.Save(&employee).Error

	return employee, err
}

func (repository *employeRepository) DeleteEmployee(employee models.Employee) (bool, error) {
	status := true
	err := repository.db.Delete(&employee).Error
	if err != nil {
		status = false
	}

	return status, err
}
