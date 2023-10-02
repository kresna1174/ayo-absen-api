package response

import "time"

type EmployeeSalaryResponse struct {
	Id         int       `json:"id"`
	CompanyId  int       `json:"company_id"`
	EmployeeId int       `json:"employee_id"`
	Salary     int       `json:"salary"`
	PayPeriod  string    `json:"pay_period"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  string    `json:"updated_by"`
}
