package entity

type User struct {
	UserId   int64  `json:"userId"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	PeriodId int64  `json:"periodId"`
	Password string `json:"password"`
}
