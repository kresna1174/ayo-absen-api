package request

type CompanyBudgetRequest struct {
	Budget    int `json:"budget" binding:"required"`
	CompanyId int `json:"company_id" binding:"required"`
	Active    int `json:"active" binding:"required"`
}
