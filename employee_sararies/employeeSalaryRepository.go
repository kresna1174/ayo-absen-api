package employee_sararies

import "gorm.io/gorm"

type EmployeeSalaryRepositoryInterface interface {
	GetAll() ([]EmployeeSalaries, error)
	FindById(Id int) (EmployeeSalaries, error)
	Create(company EmployeeSalaries) (EmployeeSalaries, error)
	Update(company EmployeeSalaries) (EmployeeSalaries, error)
	Delete(Company EmployeeSalaries) (bool, error)
}

type employeeRepositorySalary struct {
	db *gorm.DB
}

func NewEmployeeRepositorySalaryRepository(db *gorm.DB) *employeeRepositorySalary {
	return &employeeRepositorySalary{db}
}

func (r *employeeRepositorySalary) GetAll() ([]EmployeeSalaries, error) {
	var employeeSalaryEntity []EmployeeSalaries
	err := r.db.Find(&employeeSalaryEntity).Error
	if err != nil {

	}
	return employeeSalaryEntity, err
}

func (r *employeeRepositorySalary) FindById(Id int) (EmployeeSalaries, error) {
	var employeeSalaryEntity EmployeeSalaries
	err := r.db.Find(&employeeSalaryEntity, Id).Error

	if err != nil {

	}
	return employeeSalaryEntity, err
}

func (r *employeeRepositorySalary) Create(employee EmployeeSalaries) (EmployeeSalaries, error) {
	err := r.db.Create(&employee).Error

	if err != nil {

	}
	return employee, err
}

func (r *employeeRepositorySalary) Update(employee EmployeeSalaries) (EmployeeSalaries, error) {
	err := r.db.Save(&employee).Error
	if err != nil {

	}
	return employee, err
}

func (r *employeeRepositorySalary) Delete(employee EmployeeSalaries) (bool, error) {
	status := true
	err := r.db.Delete(&employee).Error
	if err != nil {
		status = false
	}
	return status, err
}
