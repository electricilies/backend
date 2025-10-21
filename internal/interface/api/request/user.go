package request

import (
	"backend/internal/domain/user"
	"time"
)

type User struct {
	ID          string     `json:"id" binding:"required"`
	FirstName   string     `json:"first_name" binding:"required"`
	LastName    string     `json:"last_name" binding:"required"`
	Email       string     `json:"email" binding:"required,email"`
	Birthday    *time.Time `json:"birthday,omitempty"` // TODO: Check if json can parse to date
	PhoneNumber string     `json:"phone_number,omitempty"`
	Address     string     `json:"address,omitempty"`
}

func (r User) ToDomain() *user.User {
	return &user.User{
		ID:          r.ID,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Email:       r.Email,
		Birthday:    r.Birthday,
		PhoneNumber: r.PhoneNumber,
		Address:     r.Address,
	}
}
