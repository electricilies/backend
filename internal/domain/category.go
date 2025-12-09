package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
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
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (c *Category) Update(name string) error {
	updated := false
	if name != "" && c.Name != name {
		c.Name = name
		updated = true
	}
	if updated {
		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(c); err != nil {
			return err
		}
		c.UpdatedAt = time.Now()
	}
	return nil
}
