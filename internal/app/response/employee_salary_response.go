package response

type EmployeeSalaryResponse struct {
	CompanyId  int    `json:"company_id" binding:"required"`
	EmployeeId int    `json:"employee_id" binding:"required"`
	Salary     int    `json:"salary" binding:"required"`
	PayPeriod  string `json:"pay_period" binding:"required"`
}
