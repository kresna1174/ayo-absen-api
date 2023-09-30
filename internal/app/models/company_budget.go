package models

import "time"

type CompanyBudget struct {
	Id        int `json:"id"`
	CompanyId int `json:"company_id"`
	Budget    int `json:"budget"`
	Active    int `json:"active"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

type Tabler interface {
	TableName() string
}

func (CompanyBudget) TableName() string {
	return "company_budget"
}
