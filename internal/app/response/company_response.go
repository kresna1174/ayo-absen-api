package response

type CompanyResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Active int    `json:"active"`
}
