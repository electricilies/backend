package domain

import (
	"time"
)

type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Username    string    `json:"userName"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	PhoneNumber string    `json:"phoneNumber"`
	Address     string    `json:"address,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
}
