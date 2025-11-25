package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	queries *sqlc.Queries
}

var _ domain.CategoryRepository = (*Category)(nil)

func ProvideCategory(q *sqlc.Queries) *Category {
	return &Category{queries: q}
}

func (r *Category) Count(ctx context.Context) (*int, error) {
	count, err := r.queries.CountCategories(ctx, sqlc.CountCategoriesParams{
		Deleted: string(domain.DeletedExcludeParam),
	})
	return ptr.To(int(count)), err
}

func (r *Category) List(
	ctx context.Context,
	ids *[]uuid.UUID,
	search *string,
	limit int,
	offset int,
) (*[]domain.Category, error) {
	categories, err := r.queries.ListCategories(ctx, sqlc.ListCategoriesParams{
		Search:  search,
		IDs:     ptr.Deref(ids, []uuid.UUID{}),
		Deleted: string(domain.DeletedExcludeParam),
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	result := make([]domain.Category, 0, len(categories))
	for _, cat := range categories {
		result = append(result, domain.Category{
			ID:        cat.ID,
			Name:      cat.Name,
			CreatedAt: cat.CreatedAt.Time,
			UpdatedAt: cat.UpdatedAt.Time,
			DeletedAt: fromPgValidToPtr(cat.DeletedAt.Time, cat.DeletedAt.Valid),
		})
	}
	return &result, nil
}

func (r *Category) Get(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	cat, err := r.queries.GetCategory(ctx, sqlc.GetCategoryParams{
		ID:      id,
		Deleted: string(domain.DeletedExcludeParam),
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	result := domain.Category{
		ID:        cat.ID,
		Name:      cat.Name,
		CreatedAt: cat.CreatedAt.Time,
		UpdatedAt: cat.UpdatedAt.Time,
		DeletedAt: fromPgValidToPtr(cat.DeletedAt.Time, cat.DeletedAt.Valid),
	}
	return &result, nil
}

func (r *Category) Save(ctx context.Context, category domain.Category) error {
	return r.queries.UpsertCategory(ctx, sqlc.UpsertCategoryParams{
		ID:   category.ID,
		Name: category.Name,
		CreatedAt: pgtype.Timestamptz{
			Time:  category.CreatedAt,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  category.UpdatedAt,
			Valid: true,
		},
		DeletedAt: pgtype.Timestamptz{
			Time:  ptr.Deref(category.DeletedAt, category.CreatedAt),
			Valid: category.DeletedAt != nil,
		},
	})
}
