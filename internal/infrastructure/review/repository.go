package review

import (
	"context"

	"backend/internal/domain/review"
	"backend/internal/infrastructure/persistence/postgres"
)

type repositoryImpl struct {
	db *postgres.Queries
}

func NewRepository(
	db *postgres.Queries,
) review.Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) ListByProduct(ctx context.Context, productID int, queryParams *review.QueryParams) (*review.Pagination, error) {
	return &review.Pagination{}, nil
}

func (r *repositoryImpl) Get(ctx context.Context, id int) (*review.Model, error)

func (r *repositoryImpl) Create(ctx context.Context, review *review.Model) (*review.Model, error)

func (r *repositoryImpl) Update(ctx context.Context, review *review.Model, id int) (*review.Model, error)

func (r *repositoryImpl) Delete(ctx context.Context, id int) error
