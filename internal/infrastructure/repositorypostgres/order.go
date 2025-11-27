package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repositorypostgres/sqlc"
)

type Order struct {
	queries *sqlc.Queries
}

var _ domain.OrderRepository = (*Order)(nil)

func ProvideOrder(q *sqlc.Queries) *Order {
	return &Order{queries: q}
}

func (r *Order) List(ctx context.Context, params domain.OrderRepositoryListParam) (*[]domain.Order, error) {
	panic("implement me")
}

func (r *Order) Count(ctx context.Context, params domain.OrderRepositoryCountParam) (*int, error) {
	panic("implement me")
}

func (r *Order) Get(ctx context.Context, params domain.OrderRepositoryGetParam) (*domain.Order, error) {
	panic("implement me")
}

func (r *Order) Save(ctx context.Context, params domain.OrderRepositorySaveParam) error {
	panic("implement me")
}
