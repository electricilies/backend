package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	queries *sqlc.Queries
}

var _ domain.CategoryRepository = (*Category)(nil)

func ProvideCategory(q *sqlc.Queries) *Category {
	return &Category{queries: q}
}

func (r *Category) Count(ctx context.Context, params domain.CategoryRepositoryCountParam) (*int, error) {
	count, err := r.queries.CountCategories(ctx, sqlc.CountCategoriesParams{
		Search:  params.Search,
		IDs:     params.IDs,
		Deleted: string(params.Deleted),
	})
	return ptr.To(int(count)), err
}

func (r *Category) List(
	ctx context.Context,
	params domain.CategoryRepositoryListParam,
) (*[]domain.Category, error) {
	categories, err := r.queries.ListCategories(ctx, sqlc.ListCategoriesParams{
		Search:  params.Search,
		IDs:     params.IDs,
		Deleted: string(params.Deleted),
		Limit:   int32(params.Limit),
		Offset:  int32(params.Offset),
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	result := make([]domain.Category, 0, len(categories))
	for _, cat := range categories {
		result = append(result, domain.Category{
			ID:        cat.ID,
			Name:      cat.Name,
			CreatedAt: cat.CreatedAt.Time,
			UpdatedAt: cat.UpdatedAt.Time,
			DeletedAt: cat.DeletedAt.Time,
		})
	}
	return &result, nil
}

func (r *Category) Get(ctx context.Context, params domain.CategoryRepositoryGetParam) (*domain.Category, error) {
	cat, err := r.queries.GetCategory(ctx, sqlc.GetCategoryParams{
		ID:      params.ID,
		Deleted: string(domain.DeletedExcludeParam),
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	result := domain.Category{
		ID:        cat.ID,
		Name:      cat.Name,
		CreatedAt: cat.CreatedAt.Time,
		UpdatedAt: cat.UpdatedAt.Time,
		DeletedAt: cat.DeletedAt.Time,
	}
	return &result, nil
}

func (r *Category) Save(ctx context.Context, params domain.CategoryRepositorySaveParam) error {
	return r.queries.UpsertCategory(ctx, sqlc.UpsertCategoryParams{
		ID:   params.Category.ID,
		Name: params.Category.Name,
		CreatedAt: pgtype.Timestamptz{
			Time:  params.Category.CreatedAt,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  params.Category.UpdatedAt,
			Valid: true,
		},
		DeletedAt: pgtype.Timestamptz{
			Time:  params.Category.DeletedAt,
			Valid: !params.Category.DeletedAt.IsZero(),
		},
	})
}
