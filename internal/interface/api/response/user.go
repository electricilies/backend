package response

import (
	"time"

	"backend/internal/domain/user"
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

func UserFromDomain(u *user.Model) *User {
	return &User{
		ID:          u.ID.String(),
		FirstName:   *u.FirstName,
		LastName:    *u.LastName,
		Username:    *u.UserName,
		Email:       *u.Email,
		DateOfBirth: *u.DateOfBirth,
		Address:     *u.Address,
		PhoneNumber: *u.PhoneNumber,
		CreatedAt:   *u.CreatedAt,
	}
}

func UsersFromDomain(users []*user.Model) []*User {
	responses := make([]*User, len(users))
	for i, u := range users {
		responses[i] = UserFromDomain(u)
	}

	return responses
}
