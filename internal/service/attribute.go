package service

import (
	"time"

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

func (a *Attribute) Update(
	attribute *domain.Attribute,
	name *string,
) error {
	if name != nil {
		attribute.Name = *name
	}
	if err := a.validate.Struct(attribute); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (a *Attribute) AddValues(attribute domain.Attribute, attributeValues ...domain.AttributeValue) error {
	if attribute.Values == nil {
		attribute.Values = &[]domain.AttributeValue{}
	}
	*attribute.Values = append(*attribute.Values, attributeValues...)
	err := a.validate.Struct(attribute)
	if err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
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

func (a *Attribute) UpdateValue(
	attribute domain.Attribute,
	attributeValueID uuid.UUID,
	value *string,
) error {
	for i, v := range *attribute.Values {
		if v.ID == attributeValueID {
			if value != nil {
				(*attribute.Values)[i].Value = *value
			}
			break
		}
	}
	if err := a.validate.Struct(attribute); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (a *Attribute) Remove(
	attribute *domain.Attribute,
) error {
	if attribute == nil {
		return domain.ErrInvalid
	}
	now := time.Now()
	attribute.DeletedAt = &now
	for i := range *attribute.Values {
		(*attribute.Values)[i].DeletedAt = &now
	}
	if err := a.validate.Struct(attribute); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (a *Attribute) RemoveValue(
	attribute domain.Attribute,
	attributeValueID uuid.UUID,
) error {
	if attribute.Values == nil {
		return nil
	}
	newValues := []domain.AttributeValue{}
	for _, v := range *attribute.Values {
		if v.ID != attributeValueID {
			newValues = append(newValues, v)
		}
	}
	attribute.Values = &newValues
	if err := a.validate.Struct(attribute); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
