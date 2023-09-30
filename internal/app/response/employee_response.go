package response

import "time"

type EmployeeResponse struct {
	Id        int       `json:"id"`
	CompanyId int       `json:"company_id"`
	Name      string    `json:"name"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Active    int       `json:"active"`
}
