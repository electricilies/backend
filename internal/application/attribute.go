package application

import (
	"context"

	"backend/internal/domain/attribute"
)

type Attribute interface {
	ListAttributes(ctx context.Context, queryParams *attribute.QueryParams) (*attribute.PaginationModel, error)
}

type attributeApp struct {
	attributeRepo attribute.Repository
}

func NewAttribute(attributeRepo attribute.Repository) Attribute {
	return &attributeApp{
		attributeRepo: attributeRepo,
	}
}

func (a *attributeApp) ListAttributes(ctx context.Context, queryParams *attribute.QueryParams) (*attribute.PaginationModel, error) {
	return a.attributeRepo.List(ctx, queryParams)
}
