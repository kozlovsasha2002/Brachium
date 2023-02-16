package entity

type CheckList struct {
	ListId                 int    `json:"listId" db:"listId"'`
	OverallPercentComplete int8   `json:"overallPercentComplete" db:"overallPercentComplete"`
	StartDate              string `json:"startDate" db:"startDate" binding:"required"`
	EndDate                string `json:"endDate" db:"endDate" binding:"required"`
	UserId                 int    `json:"userId" db:"userId"`
}

type UpdateCheckListInput struct {
	OverallPercentComplete *int8   `json:"overallPercentComplete" db:"overallPercentComplete"`
	StartDate              *string `json:"startDate" db:"startDate" binding:"required"`
	EndDate                *string `json:"endDate" db:"endDate" binding:"required"`
}
