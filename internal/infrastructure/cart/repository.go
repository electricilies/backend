package cart

import (
	"context"

	"backend/internal/domain/cart"
	"backend/internal/infrastructure/mapper"
	"backend/internal/infrastructure/persistence/postgres"

	"github.com/jackc/pgx/v5/pgtype"
)

type RepositoryImpl struct {
	db *postgres.Queries
}

func NewRepository() cart.Repository {
	return &RepositoryImpl{}
}

func ProvideRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Get(ctx context.Context, id int, queryParams *cart.QueryParams) (*cart.Model, error) {
	cartEntity, err := r.db.GetCart(ctx, postgres.GetCartParams{
		ID: pgtype.Int4{
			Int32: int32(id),
			Valid: true,
		},
	})
	if err != nil {
		return nil, mapper.ToDomainErrorFromPostgres(err)
	}
	return ToDomain(&cartEntity), nil
}

func (r *RepositoryImpl) AddItem(ctx context.Context, itemModel *cart.ItemModel) (*cart.ItemModel, error) {
	return nil, nil
}

func (r *RepositoryImpl) UpdateItem(
	ctx context.Context,
	itemModel *cart.ItemModel,
	id string,
) (*cart.ItemModel, error) {
	return &cart.ItemModel{}, nil
}

func (r *RepositoryImpl) RemoveItem(ctx context.Context, id string) error {
	return nil
}
