package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	Id        int       `json:"id"`
	CompanyId int       `json:"company_id"`
	Name      string    `json:"name"`
	Start     string    `json:"start"`
	End       string    `json:"end"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type EmployeeWithCompany struct {
	Employee
	CompanyName string `json:"company_name"`
}

func (e *Employee) ViewCompany(db *gorm.DB) ([]EmployeeWithCompany, error) {
	var employeesWithCompanies []EmployeeWithCompany

	err := db.Model(&e).
		Select("employees.*, companies.name AS company_name").
		Joins("JOIN companies ON companies.id = employees.company_id").
		Scan(&employeesWithCompanies).
		Error

	if err != nil {
		return nil, err
	}

	return employeesWithCompanies, nil
}

func (e *Employee) ViewCompanySinggle(db *gorm.DB, Id int) (EmployeeWithCompany, error) {
	var employeesWithCompanies EmployeeWithCompany

	err := db.Model(&e).
		Select("employees.*, companies.name AS company_name").
		Joins("JOIN companies ON companies.id = employees.company_id").
		Where("employees.id = ?", Id).
		Scan(&employeesWithCompanies).
		Error

	if err != nil {
		return EmployeeWithCompany{}, err
	}

	return employeesWithCompanies, nil
}
