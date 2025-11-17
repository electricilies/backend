package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/jackc/pgx/v5/pgtype"
)

type PostgresAttribute struct {
	db *postgres.Queries
}

func ProvidePostgresAttribute(db *postgres.Queries) *PostgresAttribute {
	return &PostgresAttribute{
		db: db,
	}
}

var _ repository.Attribute = &PostgresAttribute{}

func (r *PostgresAttribute) GetAttribute(
	ctx context.Context,
	param service.GetAttributeParam,
) (*domain.Attribute, error) {
	attribute, err := r.db.GetAttribute(ctx,
		postgres.GetAttributeParams{
			ID: int32(param.AttributeID),
		},
	)
	if err != nil {
		return nil, err
	}

	return &domain.Attribute{
		ID:              int(attribute.ID),
		Code:            attribute.Code,
		Name:            attribute.Name,
		AttributeValues: []domain.AttributeValue{},
		DeletedAt: mapPgtypeToDomainPtr(
			attribute.DeletedAt.Time,
			attribute.DeletedAt.Valid,
		),
	}, nil
}

func (r *PostgresAttribute) ListAttributes(
	ctx context.Context,
	param service.ListAttributesParam,
) (*domain.Pagination[domain.Attribute], error) {
	search, searchValid := domainPtrToPgtype(param.Search)
	ids, _ := domainPtrSliceToPgtype(param.AttributeIDs, func(id int) int32 {
		return int32(id)
	})
	attributes, err := r.db.ListAttributes(ctx, postgres.ListAttributesParams{
		Search: pgtype.Text{
			String: search,
			Valid:  searchValid,
		},
		IDs: ids,
		Limit: pgtype.Int4{
			Int32: int32(param.Limit),
			Valid: true,
		},
		Offset: pgtype.Int4{
			Int32: int32(param.Page * param.Limit),
			Valid: true,
		},
		Deleted: param.Deleted,
	})
	if err != nil {
		return nil, err
	}

	if len(attributes) == 0 {
		return &domain.Pagination[domain.Attribute]{
			Meta: domain.PaginationMeta{
				CurrentPage:  param.Page,
				ItemsPerPage: param.Limit,
			},
		}, nil
	}

	data := make([]domain.Attribute, 0, len(attributes))
	for _, attribute := range attributes {
		data = append(data, domain.Attribute{
			ID:   int(attribute.ID),
			Code: attribute.Code,
			Name: attribute.Name,
			DeletedAt: mapPgtypeToDomainPtr(
				attribute.DeletedAt.Time,
				attribute.DeletedAt.Valid,
			),
		})
	}

	return &domain.Pagination[domain.Attribute]{
		Data: data,
		Meta: domain.PaginationMeta{
			TotalItems:   int(attributes[0].TotalCount),
			CurrentPage:  param.Page,
			ItemsPerPage: param.Limit,
			PageItems:    int(attributes[0].CurrentCount),
		},
	}, nil
}

func (r *PostgresAttribute) CreateAttribute(
	ctx context.Context,
	param service.CreateAttributeParam,
) (*domain.Attribute, error) {
	return nil, nil
}

func (r *PostgresAttribute) UpdateAttribute(
	ctx context.Context,
	param service.UpdateAttributeParam,
) (*domain.Attribute, error) {
	return nil, nil
}

func (r *PostgresAttribute) CreateAttributeValue(
	ctx context.Context,
	param service.CreateAttributeValueParam,
) (*domain.AttributeValue, error) {
	return nil, nil
}

func (r *PostgresAttribute) UpdateAttributeValues(
	ctx context.Context,
	param []service.UpdateAttributeValueParam,
) error {
	return nil
}
