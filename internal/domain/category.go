package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID  `json:"id"        binding:"required"                    validate:"required"`
	Name      string     `json:"name"      binding:"required"                    validate:"required,gte=2,lte=100"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"                    validate:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"                    validate:"required,gtefield=CreatedAt"`
DeletedAt *time.Time `json:"deletedAt" validate:"omitempty,gtefield=CreatedAt"`
}

func NewCategory(name string) (Category, error) {
	now := time.Now()
	id, err := uuid.NewV7()
	if err != nil {
		return Category{}, err
	}
	category := Category{
		ID:        id,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return category, nil
}

func (c *Category) Update(name *string) {
	if c == nil {
		return
	}
	updated := false
	if name != nil {
		c.Name = *name
		updated = true
	}
	if updated {
		c.UpdatedAt = time.Now()
	}
}
