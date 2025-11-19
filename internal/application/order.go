package application

import (
	"context"

	"backend/internal/domain"
)

type Order interface {
	List(ctx context.Context, param ListOrderParam) (*Pagination[domain.Order], error)
}
