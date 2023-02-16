package entity

type User struct {
	UserId   int    `json:"-" db:"id"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
