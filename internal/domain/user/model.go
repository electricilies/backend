package user

import "time"

type User struct {
	ID          string
	UserName    string
	FirstName   string
	LastName    string
	Email       string
	Address     string
	Birthday    *time.Time
	PhoneNumber string
	CreatedAt   *time.Time
}
