package request

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Active   int    `json:"active" binding:"required"`
}
