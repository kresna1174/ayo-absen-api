package response

import "time"

type CompanyBudgetResponse struct {
	Id        int       `json:"id"`
	CompanyId int       `json:"company_id"`
	Budget    int       `json:"budget"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
