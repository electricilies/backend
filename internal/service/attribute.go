package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
)

type Attribute struct {
	validate *validator.Validate
}

func ProvideAttribute(
	validate *validator.Validate,
) *Attribute {
	return &Attribute{
		validate: validate,
	}
}

var _ domain.AttributeService = &Attribute{}

func (a *Attribute) Validate(
	attribute domain.Attribute,
) error {
	if err := a.validate.Struct(attribute); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
