package models

import "time"

type Users struct {
	Id        int
	Username  string
	Name      string
	Password  string
	Email     string
	Active    int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
