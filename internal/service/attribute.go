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

var _ domain.AttributeService = (*Attribute)(nil)

func (a *Attribute) Validate(
	attribute domain.Attribute,
) error {
	if err := a.validate.Struct(attribute); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (a *Attribute) FilterAttributeValuesFromAttributes(
	attributes []domain.Attribute,
	attributeValueIDs []uuid.UUID,
) []domain.AttributeValue {
	attributeValueIDSet := make(map[uuid.UUID]struct{}, len(attributeValueIDs))
	for _, id := range attributeValueIDs {
		attributeValueIDSet[id] = struct{}{}
	}
	result := []domain.AttributeValue{}
	for _, attribute := range attributes {
		for _, value := range attribute.Values {
			if _, exists := attributeValueIDSet[value.ID]; exists {
				result = append(result, value)
			}
		}
	}
	return result
}
