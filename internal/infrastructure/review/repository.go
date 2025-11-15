package review

import (
	"context"

	"backend/internal/domain/review"
	"backend/internal/infrastructure/persistence/postgres"
)

type RepositoryImpl struct {
	db *postgres.Queries
}

func NewRepository(
	db *postgres.Queries,
) review.Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func ProvideRepository(
	db *postgres.Queries,
) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) ListByProduct(ctx context.Context, productID int, reviewQueryParams *review.QueryParams) (*review.Pagination, error) {
	return &review.Pagination{}, nil
}

func (r *RepositoryImpl) Get(ctx context.Context, id int) (*review.Model, error) {
	return nil, nil
}

func (r *RepositoryImpl) Create(ctx context.Context, reviewModel *review.Model) (*review.Model, error) {
	return nil, nil
}

func (r *RepositoryImpl) Update(ctx context.Context, reviewModel *review.Model, id int) (*review.Model, error) {
	return nil, nil
}

func (r *RepositoryImpl) Delete(ctx context.Context, id int) error {
	return nil
}
