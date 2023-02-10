package entity

type Period struct {
	PeriodId               int64  `json:"periodId"`
	TaskId                 int64  `json:"taskId"`
	OverallPercentComplete int8   `json:"overallPercentComplete"`
	StartDate              string `json:"startDate"`
	EndDate                string `json:"endDate"`
	AmountOfDays           string `json:"amountOfDays"`
}
