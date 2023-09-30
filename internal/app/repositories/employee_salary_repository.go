package repositories

import (
	"api-ayo-absen/internal/app/models"
	"gorm.io/gorm"
)

type EmployeeSalaryRepositoryInterface interface {
	GetAll() ([]models.EmployeeSalaries, error)
	FindById(Id int) (models.EmployeeSalaries, error)
	Create(company models.EmployeeSalaries) (models.EmployeeSalaries, error)
	Update(company models.EmployeeSalaries) (models.EmployeeSalaries, error)
	Delete(Company models.EmployeeSalaries) (bool, error)
}

type employeeSalaryRepository struct {
	db *gorm.DB
}

func NewEmployeeSalaryRepository(db *gorm.DB) *employeeSalaryRepository {
	return &employeeSalaryRepository{db}
}

func (r *employeeSalaryRepository) GetAll() ([]models.EmployeeSalaries, error) {
	var employeeSalaryEntity []models.EmployeeSalaries
	err := r.db.Find(&employeeSalaryEntity).Error
	if err != nil {

	}
	return employeeSalaryEntity, err
}

func (r *employeeSalaryRepository) FindById(Id int) (models.EmployeeSalaries, error) {
	var employeeSalaryEntity models.EmployeeSalaries
	err := r.db.Find(&employeeSalaryEntity, Id).Error

	if err != nil {

	}
	return employeeSalaryEntity, err
}

func (r *employeeSalaryRepository) Create(employee models.EmployeeSalaries) (models.EmployeeSalaries, error) {
	err := r.db.Create(&employee).Error

	if err != nil {

	}
	return employee, err
}

func (r *employeeSalaryRepository) Update(employee models.EmployeeSalaries) (models.EmployeeSalaries, error) {
	err := r.db.Save(&employee).Error
	if err != nil {

	}
	return employee, err
}

func (r *employeeSalaryRepository) Delete(employee models.EmployeeSalaries) (bool, error) {
	status := true
	err := r.db.Delete(&employee).Error
	if err != nil {
		status = false
	}
	return status, err
}
