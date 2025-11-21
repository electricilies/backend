package domain

import "github.com/google/uuid"

type AttributeService interface {
	Create(
		code string,
		name string,
	) (*Attribute, error)

	Update(
		attribute *Attribute,
		name *string,
	) error

	AddValues(
		attribute Attribute,
		attributeValues ...AttributeValue,
	) error

	CreateValue(
		value string,
	) (*AttributeValue, error)

	UpdateValue(
		attribute Attribute,
		attributeValueID uuid.UUID,
		value *string,
	) error

	DeleteValue(
		attribute Attribute,
		attributeValueID uuid.UUID,
	) error
}
