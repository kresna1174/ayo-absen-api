package models

import "time"

type Employee struct {
	Id        int
	CompanyId int
	Name      string
	Start     string
	End       string
	Active    int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
