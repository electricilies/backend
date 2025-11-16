package response

import (
	"time"

	"backend/internal/domain/category"
)

type Category struct {
	ID        int        `json:"id" binding:"required"`
	Name      string     `json:"name" binding:"required"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt" binding:"omitnil"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"`
}

func CategoryFromDomain(c *category.Model) *Category {
	return &Category{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		DeletedAt: c.DeletedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
