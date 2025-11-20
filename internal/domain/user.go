package domain

import (
	"time"
)

type User struct {
	ID          string    `json:"id" binding:"required" validate:"required,uuid"`
	FirstName   string    `json:"firstName" binding:"required" validate:"required"`
	LastName    string    `json:"lastName" binding:"required" validate:"required"`
	Username    string    `json:"userName" binding:"required" validate:"required"`
	Email       string    `json:"email" binding:"required" validate:"required,email"`
	DateOfBirth time.Time `json:"dateOfBirth" binding:"required" validate:"required"`
	PhoneNumber string    `json:"phoneNumber" binding:"required" validate:"required,startswith=0,min=10,max=11"`
	Address     string    `json:"address" binding:"required" validate:"required"`
	CreatedAt   time.Time `json:"createdAt" binding:"required" validate:"required"`
}
