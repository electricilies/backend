package application

import (
	"context"

	"backend/internal/domain/attribute"
)

type Attribute interface {
	ListAttributes(ctx context.Context, queryParams *attribute.QueryParams) (*attribute.PaginationModel, error)
}

type AttributeApp struct {
	attributeRepo attribute.Repository
}

func NewAttribute(attributeRepo attribute.Repository) Attribute {
	return &AttributeApp{
		attributeRepo: attributeRepo,
	}
}

func ProvideAttribute(
	attributeRepo attribute.Repository,
) *AttributeApp {
	return &AttributeApp{
		attributeRepo: attributeRepo,
	}
}

func (a *AttributeApp) ListAttributes(ctx context.Context, queryParams *attribute.QueryParams) (*attribute.PaginationModel, error) {
	return a.attributeRepo.List(ctx, queryParams)
}
