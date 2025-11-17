package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/service"
)

type Attribute interface {
	GetAttribute(context.Context, service.GetAttributeParam) (*domain.Attribute, error)
	ListAttributes(context.Context, service.ListAttributesParam) (*domain.Pagination[domain.Attribute], error)
	CreateAttribute(context.Context, service.CreateAttributeParam) (*domain.Attribute, error)
	UpdateAttribute(context.Context, service.UpdateAttributeParam) (*domain.Attribute, error)
	CreateAttributeValue(context.Context, service.CreateAttributeValueParam) (*domain.AttributeValue, error)
	UpdateAttributeValues(context.Context, []service.UpdateAttributeValueParam) error
}
