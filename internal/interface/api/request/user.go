package request

import (
	"time"

	"backend/internal/domain/user"

	"github.com/google/uuid"
)

type CreateUser struct {
	ID string `json:"id" binding:"required"`
}

func (r *CreateUser) ToDomain() *user.Model {
	return &user.Model{
		ID: uuid.MustParse(r.ID),
	}
}

type UpdateUser struct {
	FirstName   string     `json:"firstName" binding:"required"`
	LastName    string     `json:"lastName" binding:"required"`
	Email       string     `json:"email" binding:"required,email"`
	DateOfBirth *time.Time `json:"dateOfBirth,omitempty"` // TODO: Check if json can parse to date
	PhoneNumber string     `json:"phoneNumber,omitempty"`
	Address     string     `json:"address,omitempty"`
}

func (r UpdateUser) ToDomain() *user.Model {
	return &user.Model{
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Email:       r.Email,
		DateOfBirth: r.DateOfBirth,
		PhoneNumber: r.PhoneNumber,
		Address:     r.Address,
	}
}
