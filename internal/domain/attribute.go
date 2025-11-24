package domain

import (
	"time"

	"github.com/google/uuid"
)

type Attribute struct {
	ID        uuid.UUID        `json:"id"        binding:"required"        validate:"required"               example:"123"`
	Code      string           `json:"code"      binding:"required"        validate:"required,gte=2,lte=50"  example:"color"`
	Name      string           `json:"name"      binding:"required"        validate:"required,gte=2,lte=100" example:"Color"`
	Values    []AttributeValue `json:"values"    validate:"omitempty,dive"`
	DeletedAt *time.Time       `json:"deletedAt"`
}

type AttributeValue struct {
	ID        uuid.UUID  `json:"id"        binding:"required"   validate:"required"               example:"1"`
	Value     string     `json:"value"     binding:"required"   validate:"required,gte=1,lte=100" example:"Red"`
	DeletedAt *time.Time `json:"deletedAt" validate:"omitempty"`
}

func NewAttribute(
	code string,
	name string,
) (*Attribute, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	attribute := &Attribute{
		ID:     id,
		Code:   code,
		Name:   name,
		Values: []AttributeValue{},
	}
	return attribute, nil
}

func NewAttributeValue(value string) (*AttributeValue, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	attributeValue := &AttributeValue{
		ID:    id,
		Value: value,
	}
	return attributeValue, nil
}

func (a *Attribute) Update(name *string) {
	if name != nil {
		a.Name = *name
	}
}

func (a *Attribute) GetValueByID(id uuid.UUID) *AttributeValue {
	if a.Values == nil {
		return nil
	}
	for _, value := range a.Values {
		if value.ID == id {
			return &value
		}
	}
	return nil
}

func (a *Attribute) AddValues(attributeValues ...AttributeValue) {
	a.Values = append(a.Values, attributeValues...)
}

func (a *Attribute) UpdateValue(
	attributeValueID uuid.UUID,
	value *string,
) error {
	for i, v := range a.Values {
		if v.ID == attributeValueID {
			if value != nil {
				a.Values[i].Value = *value
			}
			return nil
		}
	}
	return ErrNotFound
}

func (a *Attribute) Remove() {
	now := time.Now()
	if a.DeletedAt == nil {
		a.DeletedAt = &now
	}
	for i := range a.Values {
		if a.Values[i].DeletedAt == nil {
			a.Values[i].DeletedAt = &now
		}
	}
}

func (a *Attribute) RemoveValue(attributeValueID uuid.UUID) error {
	if a == nil {
		return ErrInvalid
	}
	newValues := []AttributeValue{}
	for _, v := range a.Values {
		if v.ID != attributeValueID {
			newValues = append(newValues, v)
		}
	}
	a.Values = newValues
	return nil
}
