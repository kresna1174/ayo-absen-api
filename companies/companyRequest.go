package companies

type CompanyRequest struct {
	Name   string `json:"name" binding:"required"`
	Active int    `json:"active" binding:"required"`
}
