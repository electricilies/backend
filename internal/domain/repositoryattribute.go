package domain

import (
	"context"

	"backend/internal/service"
)

type AttributeRepository interface {
	Get(context.Context, service.GetAttributeParam) (Attribute, error)
	List(context.Context, service.ListAttributesParam) (Pagination[Attribute], error)
	Create(context.Context, service.CreateAttributeParam) (Attribute, error)
	Update(context.Context, service.UpdateAttributeParam) (Attribute, error)
	CreateValue(context.Context, service.CreateAttributeValueParam) (AttributeValue, error)
	UpdateValues(context.Context, []service.UpdateAttributeValueParam) error
}
