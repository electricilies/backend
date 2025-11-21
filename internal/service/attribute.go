package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

func (a *Attribute) Create(
	code string,
	name string,
) (*domain.Attribute, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	attribute := &domain.Attribute{
		ID:   id,
		Code: code,
		Name: name,
	}
	if err := a.validate.Struct(attribute); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return attribute, nil
}

func (a *Attribute) CreateValue(
	value string,
) (*domain.AttributeValue, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	attributeValue := domain.AttributeValue{
		ID:    id,
		Value: value,
	}
	if err := a.validate.Struct(attributeValue); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return &attributeValue, nil
}
