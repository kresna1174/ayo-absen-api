package models

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeSalaries struct {
	Id         int       `json:"id"`
	CompanyId  int       `json:"company_id"`
	EmployeeId int       `json:"employee_id"`
	Salary     int       `json:"salary"`
	PayPeriod  string    `json:"pay_period"`
	Created_at time.Time `json:"created_at"`
	Created_by string    `json:"created_by"`
	Updated_at time.Time `json:"updated_at"`
	Updated_by string    `json:"updated_by"`
}

type EmployeeSalaryWithEmployeeAndCompany struct {
	EmployeeSalaries
	EmployeeName string `json:"employee_name"`
	CompanyName  string `json:"company_name"`
}

func (e *EmployeeSalaries) GetRelation(db *gorm.DB) ([]EmployeeSalaryWithEmployeeAndCompany, error) {
	var model []EmployeeSalaryWithEmployeeAndCompany

	err := db.Model(&e).
		Select("employee_salaries.*, employees.name as employee_name, companies.name as company_name").
		Joins("join employees on employees.id = employee_salaries.employee_id").
		Joins("join companies on companies.id = employee_salaries.company_id").
		Scan(&model).
		Error

	if err != nil {
		return []EmployeeSalaryWithEmployeeAndCompany{}, err
	}
	return model, err
}
