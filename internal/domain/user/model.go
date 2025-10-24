package user

import "time"

type User struct {
	ID          string
	UserName    string
	FirstName   string
	LastName    string
	Email       string
	Address     string
	DateOfBirth *time.Time
	PhoneNumber string
	CreatedAt   *time.Time
}
