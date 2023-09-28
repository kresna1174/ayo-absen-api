package employee

type EmployeeRequest struct {
	CompanyId int    `json:"company_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	//Start     time.Time `json:"start" binding:"required"`
	//End       time.Time `json:"end" binding:"required"`
	Active int `json:"active" binding:"required"`
}
