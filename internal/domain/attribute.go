package domain

import (
	"time"

	"github.com/google/uuid"
)

type Attribute struct {
	ID        uuid.UUID        `validate:"required"                               example:"123"`
	Code      string           `validate:"required,gte=2,lte=50"                  example:"color"`
	Name      string           `validate:"required,gte=2,lte=100"                 example:"Color"`
	Values    []AttributeValue `validate:"omitempty,uniqueAttributeValues,dive"`
	DeletedAt time.Time
}

type AttributeValue struct {
	ID        uuid.UUID `validate:"required"               example:"1"`
	Value     string    `validate:"required,gte=1,lte=100" example:"Red"`
	DeletedAt time.Time
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

func (a *Attribute) Update(name string) {
	if name != "" && a.Name != name {
		a.Name = name
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
	value string,
) error {
	for i, v := range a.Values {
		if v.ID == attributeValueID {
			if value != "" {
				a.Values[i].Value = value
			}
			return nil
		}
	}
	return ErrNotFound
}

func (a *Attribute) Remove() {
	now := time.Now()
	a.DeletedAt = now
	for i := range a.Values {
		a.Values[i].DeletedAt = now
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
