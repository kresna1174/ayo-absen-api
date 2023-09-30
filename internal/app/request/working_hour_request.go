package request

type WorkingHoursRequest struct {
	CompanyId int    `json:"company_id" binding:"required"`
	StartDay  string `json:"start_day" binding:"required"`
	EndDay    string `json:"end_day" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
	Active    int    `json:"active" binding:"required"`
}
