package response

import (
	"backend/internal/domain/user"
	"time"
)

type User struct {
	ID          string     `json:"id" binding:"required"`
	FirstName   string     `json:"first_name" binding:"required"`
	LastName    string     `json:"last_name" binding:"required"`
	Username    string     `json:"user_name" binding:"required"`
	Email       string     `json:"email" binding:"required"`
	Birthday    *time.Time `json:"birthday" binding:"required"`
	PhoneNumber string     `json:"phone_number" binding:"required"`
	Address     string     `json:"address,omitempty"`
	CreatedAt   *time.Time `json:"created_at" binding:"required"`
}

func UserFromDomain(u *user.User) *User {
	return &User{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Username:    u.UserName,
		Birthday:    u.Birthday,
		Address:     u.Address,
		PhoneNumber: u.PhoneNumber,
		CreatedAt:   u.CreatedAt,
	}
}
