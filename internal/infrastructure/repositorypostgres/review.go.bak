package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repositorypostgres/sqlc"
)

type Review struct {
	queries *sqlc.Queries
}

var _ domain.ReviewRepository = (*Review)(nil)

func ProvideReview(q *sqlc.Queries) *Review {
	return &Review{queries: q}
}

func (r *Review) List(
	ctx context.Context,
	params domain.ReviewRepositoryListParam,
) (*[]domain.Review, error) {
	panic("implement me")
}

func (r *Review) Count(ctx context.Context, params domain.ReviewRepositoryCountParam) (*int, error) {
	panic("implement me")
}

func (r *Review) Get(ctx context.Context, params domain.ReviewRepositoryGetParam) (*domain.Review, error) {
	panic("implement me")
}

func (r *Review) Save(ctx context.Context, params domain.ReviewRepositorySaveParam) error {
	panic("implement me")
}
