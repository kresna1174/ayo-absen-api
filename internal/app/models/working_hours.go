package models

import "time"

type WorkingHours struct {
	Id        int    `json:"id"`
	CompanyId int    `json:"company_id"`
	StartDay  string `json:"start_day"`
	EndDay    string `json:"end_day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Active    int    `json:"active"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
