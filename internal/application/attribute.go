package application

import (
	"context"

	"backend/internal/domain/attribute"
)

type Attribute interface {
	ListAttributes(context.Context, *attribute.QueryParams) (*attribute.PaginationModel, error)
}

type AttributeImpl struct {
	attributeRepo attribute.Repository
}

func NewAttribute(attributeRepo attribute.Repository) Attribute {
	return &AttributeImpl{
		attributeRepo: attributeRepo,
	}
}

func ProvideAttribute(
	attributeRepo attribute.Repository,
) *AttributeImpl {
	return &AttributeImpl{
		attributeRepo: attributeRepo,
	}
}

func (a *AttributeImpl) ListAttributes(
	ctx context.Context,
	queryParams *attribute.QueryParams,
) (*attribute.PaginationModel, error) {
	return a.attributeRepo.List(ctx, queryParams)
}
