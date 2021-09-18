package jsonview

type Statistics struct {
	Total_assignments             int     `json:"total_assignments"`
	Total_assignments_successful  int     `json:"total_assignments_successful"`
	Total_assignments_failed      int     `json:"total_assignments_failed"`
	Average_investment_successful float64 `json:"average_investment_successful"`
	Average_investment_failed     float64 `json:"average_investment_failed"`
}
