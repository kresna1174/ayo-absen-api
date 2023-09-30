package models

import "time"

type Employee struct {
	Id        int
	CompanyId int
	Name      string
	Start     time.Time
	End       time.Time
	Active    int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
