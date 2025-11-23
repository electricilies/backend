package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostgresCategory struct {
	queries *postgres.Queries
}

var _ domain.CategoryRepository = (*PostgresCategory)(nil)

func ProvidePostgresCategory(q *postgres.Queries) *PostgresCategory {
	return &PostgresCategory{queries: q}
}

func (r *PostgresCategory) Count(ctx context.Context) (*int, error) {
	count, err := r.queries.CountCategories(ctx, postgres.CountCategoriesParams{
		Deleted: string(domain.DeletedExcludeParam),
	})
	return ptr.To(int(count)), err
}

func (r *PostgresCategory) List(ctx context.Context, search *string, limit int, offset int) (*[]domain.Category, error) {
	categories, err := r.queries.ListCategories(ctx, postgres.ListCategoriesParams{
		Search:  search,
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

func (r *PostgresCategory) Get(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	cat, err := r.queries.GetCategory(ctx, postgres.GetCategoryParams{
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

func (r *PostgresCategory) Save(ctx context.Context, category domain.Category) error {
	return r.queries.UpsertCategory(ctx, postgres.UpsertCategoryParams{
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
