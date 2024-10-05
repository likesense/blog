package models

type User struct {
	UserID   int    `json:"-" db:"user_id"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Avatar   string `json:"avatar,omitempty" db:"avatar"`
}
