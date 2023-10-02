package request

import "time"

type EmployeeSalaryRequest struct {
	CompanyId  int       `json:"company_id" binding:"required"`
	EmployeeId int       `json:"employee_id" binding:"required"`
	Salary     int       `json:"salary" binding:"required"`
	PayPeriod  string    `json:"pay_period" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
}

type EmployeeSalaryUpdateRequest struct {
	CompanyId  int       `json:"company_id" binding:"required"`
	EmployeeId int       `json:"employee_id" binding:"required"`
	Salary     int       `json:"salary" binding:"required"`
	PayPeriod  string    `json:"pay_period" binding:"required"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  string    `json:"updated_by"`
}
