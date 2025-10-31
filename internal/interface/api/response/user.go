package response

import (
	"backend/internal/domain/user"
	"time"
)

type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Username    string    `json:"user_name"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

func UserFromDomain(u *user.User) *User {
	return &User{
		ID:          u.ID.String(),
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Username:    u.UserName,
		Email:       u.Email,
		DateOfBirth: u.DateOfBirth,
		Address:     u.Address,
		PhoneNumber: u.PhoneNumber,
		CreatedAt:   u.CreatedAt,
	}
}

func UsersFromDomain(users []*user.User) []*User {
	responses := make([]*User, len(users))
	for i, u := range users {
		responses[i] = UserFromDomain(u)
	}

	return responses
}
