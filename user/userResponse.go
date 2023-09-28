package user

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Active   int    `json:"active"`
}
