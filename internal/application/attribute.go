package application

import (
	"context"

	"backend/internal/domain"
)

type Attribute interface {
	List(ctx context.Context, param ListAttributesParam) (*Pagination[domain.Attribute], error)
	ListValues(ctx context.Context, param ListAttributeValuesParam) (*Pagination[domain.AttributeValue], error)
}
