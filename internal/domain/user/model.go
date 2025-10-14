package user

import "time"

type User struct {
	ID          string
	Avatar      string
	FirstName   string
	LastName    string
	Username    string
	Email       string
	Birthday    *time.Time
	PhoneNumber string
	CreatedAt   *time.Time
	DeletedAt   *time.Time
}
