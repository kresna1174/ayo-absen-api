package request

import "time"

type CompanyRequest struct {
	Name      string    `json:"name" binding:"required"`
	Active    int       `json:"active" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

type CompanyUpdateRequest struct {
	Name      string    `json:"name" binding:"required"`
	Active    int       `json:"active" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
