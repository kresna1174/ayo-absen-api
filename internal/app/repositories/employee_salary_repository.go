package repositories

import (
	"api-ayo-absen/internal/app/models"
	"errors"

	"gorm.io/gorm"
)

type EmployeeSalaryRepositoryInterface interface {
	GetAll() ([]models.EmployeeSalaryWithEmployeeAndCompany, error)
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

func (r *employeeSalaryRepository) GetAll() ([]models.EmployeeSalaryWithEmployeeAndCompany, error) {
	var employeeSalaryEntity models.EmployeeSalaries
	return employeeSalaryEntity.GetRelation(r.db)
	// err := r.db.Find(&employeeSalaryEntity).Error
	// if err != nil {
	// 	return []models.EmployeeSalaryWithEmployeeAndCompany{}, err
	// }
}

func (r *employeeSalaryRepository) FindById(Id int) (models.EmployeeSalaries, error) {
	var employeeSalaryEntity models.EmployeeSalaries
	err := r.db.Find(&employeeSalaryEntity, Id).Error

	if err != nil {
		return models.EmployeeSalaries{}, err
	}
	return employeeSalaryEntity, err
}

func (r *employeeSalaryRepository) Create(employee models.EmployeeSalaries) (models.EmployeeSalaries, error) {
	err := r.db.Create(&employee).Error

	if err != nil {
		return models.EmployeeSalaries{}, err
	}
	return employee, err
}

func (r *employeeSalaryRepository) Update(employee models.EmployeeSalaries) (models.EmployeeSalaries, error) {
	err := r.db.Model(&employee).Updates(&employee).Error
	var employeeModel models.EmployeeSalaries
	er := r.db.Find(&employeeModel, employee.Id).Error
	if er != nil {
		return models.EmployeeSalaries{}, errors.New("tidak dapat mengembalikan nilai")
	}
	return employeeModel, err
}

func (r *employeeSalaryRepository) Delete(employee models.EmployeeSalaries) (bool, error) {
	status := true
	err := r.db.Delete(&employee).Error
	if err != nil {
		status = false
	}
	return status, err
}
