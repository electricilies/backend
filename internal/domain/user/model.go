package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	UserName    string
	FirstName   string
	LastName    string
	Email       string
	Address     string
	DateOfBirth *time.Time
	PhoneNumber string
	CreatedAt   *time.Time
}
