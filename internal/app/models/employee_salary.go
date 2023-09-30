package models

import "time"

type EmployeeSalaries struct {
	Id         int
	CompanyId  int
	EmployeeId int
	Salary     int
	PayPeriod  string
	Created_at time.Time
	Created_by string
	Updated_at time.Time
	Updated_by string
}
