package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `validate:"required"`
	Name      string    `validate:"required,gte=2,lte=100"`
	CreatedAt time.Time `validate:"required"`
	UpdatedAt time.Time `validate:"required,gtefield=CreatedAt"`
	DeletedAt time.Time `validate:"omitempty,gtefield=CreatedAt"`
}

func NewCategory(name string) (*Category, error) {
	now := time.Now()
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	category := &Category{
		ID:        id,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return category, nil
}

func (c *Category) Update(name string) error {
	if name != "" && c.Name != name {
		c.Name = name
		c.UpdatedAt = time.Now()
	}
	return nil
}
