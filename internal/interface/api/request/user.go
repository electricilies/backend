package request

import (
	"backend/internal/domain/user"
	"time"

	"github.com/google/uuid"
)

type CreateUser struct {
	ID string `json:"id" binding:"required"`
}

func (r *CreateUser) ToDomain() *user.User {
	return &user.User{
		ID: uuid.MustParse(r.ID),
	}
}

type UpdateUser struct {
	FirstName   string     `json:"first_name" binding:"required"`
	LastName    string     `json:"last_name" binding:"required"`
	Email       string     `json:"email" binding:"required,email"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"` // TODO: Check if json can parse to date
	PhoneNumber string     `json:"phone_number,omitempty"`
	Address     string     `json:"address,omitempty"`
}

func (r UpdateUser) ToDomain() *user.User {
	return &user.User{
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Email:       r.Email,
		DateOfBirth: r.DateOfBirth,
		PhoneNumber: r.PhoneNumber,
		Address:     r.Address,
	}
}
