package domain

import "github.com/google/uuid"

type AttributeService interface {
	Validate(
		attribute Attribute,
	) error

	FilterAttributeValuesFromAttributes(
		attributes []Attribute,
		attributeValueIDs []uuid.UUID,
	) []AttributeValue
}
