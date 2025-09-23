package user

import (
	"backend/internal/domain/user"
	"backend/internal/infrastructure/presistence/postgres"
)

func ToDomain(u postgres.User) *user.User {
	return &user.User{
		ID:   u.ID.String(),
		Name: u.Name,
	}
}
