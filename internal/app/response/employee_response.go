package response

import "time"

type EmployeeResponse struct {
	Id        int       `json:"id"`
	CompanyId int       `json:"company_id"`
	Name      string    `json:"name"`
	Start     string    `json:"start"`
	End       string    `json:"end"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
