package repositoryimpl

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/jackc/pgx/v5/pgtype"
)

type PostgresCategory struct {
	db *postgres.Queries
}

func ProvidePostgresCategory(db *postgres.Queries) *PostgresCategory {
	return &PostgresCategory{
		db: db,
	}
}

var _ repository.Category = &PostgresCategory{}

func (r *PostgresCategory) List(
	ctx context.Context,
	param service.ListCategoryParam,
) (*domain.Pagination[domain.Category], error) {
	categories, err := r.db.ListCategories(ctx, postgres.ListCategoriesParams{
		Search: pgtype.Text{
			Valid: false,
		},
		Limit: pgtype.Int4{
			Int32: int32(param.Limit),
			Valid: true,
		},
		Offset: pgtype.Int4{
			Int32: int32(param.Page * param.Limit),
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return &domain.Pagination[domain.Category]{
			Meta: domain.PaginationMeta{
				CurrentPage:  param.Page,
				ItemsPerPage: param.Limit,
			},
		}, nil
	}

	data := make([]domain.Category, 0, len(categories))
	for _, category := range categories {
		data = append(data, mapListCategoryRowToDomain(category))
	}

	return &domain.Pagination[domain.Category]{
		Data: data,
		Meta: buildPaginationMeta(
			param.Page,
			param.Limit,
			int(categories[0].TotalCount),
			int(categories[0].CurrentCount),
		),
	}, nil
}

func (r *PostgresCategory) Get(
	ctx context.Context,
	param service.GetCategoryParam,
) (*domain.Category, error) {
	category, err := r.db.GetCategory(ctx,
		postgres.GetCategoryParams{
			ID: int32(param.CategoryID),
		},
	)
	if err != nil {
		return nil, err
	}

	return mapPostgresCategoryToDomain(category), nil
}

func (r *PostgresCategory) Create(
	ctx context.Context,
	param service.CreateCategoryParam,
) (*domain.Category, error) {
	category, err := r.db.CreateCategory(ctx,
		postgres.CreateCategoryParams{
			Name: param.Name,
		},
	)
	if err != nil {
		return nil, err
	}

	return mapPostgresCategoryToDomain(category), nil
}

func (r *PostgresCategory) Update(
	ctx context.Context,
	param service.UpdateCategoryParam,
) (*domain.Category, error) {
	category, err := r.db.UpdateCategory(ctx,
		postgres.UpdateCategoryParams{
			ID:   int32(param.CategoryID),
			Name: param.Name,
			UpdatedAt: pgtype.Timestamp{
				Valid: false,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return mapPostgresCategoryToDomain(category), nil
}

// Helper functions

func mapPostgresCategoryToDomain(cat postgres.Category) *domain.Category {
	return &domain.Category{
		ID:   int(cat.ID),
		Name: cat.Name,
		CreatedAt: mapPgtypeToDomain(
			cat.CreatedAt.Time,
			cat.CreatedAt.Valid,
		),
		UpdatedAt: mapPgtypeToDomain(
			cat.UpdatedAt.Time,
			cat.UpdatedAt.Valid,
		),
		DeletedAt: mapPgtypeToDomainPtr(
			cat.DeletedAt.Time,
			cat.DeletedAt.Valid,
		),
	}
}

func mapListCategoryRowToDomain(cat postgres.ListCategoriesRow) domain.Category {
	return domain.Category{
		ID:   int(cat.ID),
		Name: cat.Name,
		CreatedAt: mapPgtypeToDomain(
			cat.CreatedAt.Time,
			cat.CreatedAt.Valid,
		),
		UpdatedAt: mapPgtypeToDomain(
			cat.UpdatedAt.Time,
			cat.UpdatedAt.Valid,
		),
		DeletedAt: mapPgtypeToDomainPtr(
			cat.DeletedAt.Time,
			cat.DeletedAt.Valid,
		),
	}
}
