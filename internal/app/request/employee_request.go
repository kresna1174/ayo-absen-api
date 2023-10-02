package request

import "time"

type EmployeeRequest struct {
	CompanyId int       `json:"company_id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Start     string    `json:"start" binding:"required"`
	End       string    `json:"end"`
	Active    int       `json:"active" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}
type EmployeeUpdateRequest struct {
	CompanyId int       `json:"company_id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Start     string    `json:"start" binding:"required"`
	End       string    `json:"end"`
	Active    int       `json:"active" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
