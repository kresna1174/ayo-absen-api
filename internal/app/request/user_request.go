package request

import "time"

type UserRequest struct {
	Username  string    `json:"username" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Active    int       `json:"active" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

type UserUpdateRequest struct {
	Username  string    `json:"username" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Password  string    `json:"password"`
	Email     string    `json:"email" binding:"required"`
	Active    int       `json:"active" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
