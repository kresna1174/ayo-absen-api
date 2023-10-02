package response

import "time"

type CompanyResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdateAt  time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
