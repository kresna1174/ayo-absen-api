package request

import "time"

type CompanyBudgetRequest struct {
	Budget    int       `json:"budget" binding:"required"`
	CompanyId int       `json:"company_id" binding:"required"`
	Active    int       `json:"active" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

type CompanyBudgetUpdateRequest struct {
	Budget    int       `json:"budget" binding:"required"`
	CompanyId int       `json:"company_id" binding:"required"`
	Active    int       `json:"active" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
