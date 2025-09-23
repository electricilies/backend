package request

import "backend/internal/domain/user"

type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name" binding:"required"`
}

func (r User) ToDomain() *user.User {
	return &user.User{
		ID:   r.ID,
		Name: r.Name,
	}
}
