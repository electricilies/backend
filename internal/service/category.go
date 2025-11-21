package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

type Category struct {
	validate *validator.Validate
}

func ProvideCategory(
	validate *validator.Validate,
) *Category {
	return &Category{
		validate: validate,
	}
}

var _ domain.CategoryService = &Category{}

func (c *Category) Create(
	name string,
) (*domain.Category, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	category := &domain.Category{
		ID:   id,
		Name: name,
	}
	if err := c.validate.Struct(category); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return category, nil
}
