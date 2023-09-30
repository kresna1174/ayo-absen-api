package response

type CompanyBudgetResponse struct {
	Id        int `json:"id"`
	CompanyId int `json:"company_id"`
	Budget    int `json:"budget"`
	Active    int `json:"active"`
}
