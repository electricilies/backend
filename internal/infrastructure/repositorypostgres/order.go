package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
)

type Order struct {
	queries *sqlc.Queries
}

var _ domain.OrderRepository = (*Order)(nil)

func ProvideOrder(q *sqlc.Queries) *Order {
	return &Order{queries: q}
}

func (r *Order) List(ctx context.Context, ids *[]uuid.UUID, search *string, deleted domain.DeletedParam, limit int, offset int) (*[]domain.Order, error) {
	panic("implement me")
}

func (r *Order) Count(ctx context.Context, ids *[]uuid.UUID, deleted domain.DeletedParam) (*int, error) {
	panic("implement me")
}

func (r *Order) Get(ctx context.Context, id uuid.UUID) (*domain.Order, error) {
	panic("implement me")
}

func (r *Order) Save(ctx context.Context, order domain.Order) error {
	panic("implement me")
}
