package response

import (
	"backend/internal/domain/user"
	"time"
)

type User struct {
	ID          string     `json:"id" binding:"required"`
	Avatar      string     `json:"avatar" binding:"required"`
	FirstName   string     `json:"first_name" binding:"required"`
	LastName    string     `json:"last_name" binding:"required"`
	Username    string     `json:"user_name" binding:"required"`
	Email       string     `json:"email" binding:"required"`
	Birthday    *time.Time `json:"birthday" binding:"required"`
	PhoneNumber string     `json:"phone_number" binding:"required"`
	CreatedAt   *time.Time `json:"created_at" binding:"required"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func UserFromDomain(u *user.User) *User {
	return &User{
		ID:          u.ID,
		Avatar:      u.Avatar,
		Birthday:    u.Birthday,
		PhoneNumber: u.PhoneNumber,
		CreatedAt:   u.CreatedAt,
		DeletedAt:   u.DeletedAt,
	}
}
