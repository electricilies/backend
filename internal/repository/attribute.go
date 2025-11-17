package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/service"
)

type Attribute interface {
	Get(context.Context, service.GetAttributeParam) (*domain.Attribute, error)
	List(context.Context, service.ListAttributesParam) (*domain.Pagination[domain.Attribute], error)
	Create(context.Context, service.CreateAttributeParam) (*domain.Attribute, error)
	Update(context.Context, service.UpdateAttributeParam) (*domain.Attribute, error)
	CreateValue(context.Context, service.CreateAttributeValueParam) (*domain.AttributeValue, error)
	UpdateValues(context.Context, []service.UpdateAttributeValueParam) error
}
