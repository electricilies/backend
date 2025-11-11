package response

import (
	"time"

	"backend/internal/domain/category"
)

type Category struct {
	ID          int       `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"createdAt" binding:"required"`
}

func CategoryFromDomain(c *category.Model) *Category {
	return &Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
	}
}
