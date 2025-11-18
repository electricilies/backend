package domain

import (
	"time"
)

type User struct {
	ID          string    `json:"id" binding:"required,uuid"`
	FirstName   string    `json:"firstName" binding:"required"`
	LastName    string    `json:"lastName" binding:"required"`
	Username    string    `json:"userName" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	DateOfBirth time.Time `json:"dateOfBirth" binding:"required"`
	PhoneNumber string    `json:"phoneNumber" binding:"required"`
	Address     string    `json:"address" binding:"required"`
	CreatedAt   time.Time `json:"createdAt" binding:"required"`
}
