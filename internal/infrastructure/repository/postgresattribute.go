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

func (r *PostgresAttribute) Get(
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

	return mapPostgresAttributeToDomain(attribute), nil
}

func (r *PostgresAttribute) List(
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
		data = append(data, mapListAttributeRowToDomain(attribute))
	}

	return &domain.Pagination[domain.Attribute]{
		Data: data,
		Meta: buildPaginationMeta(
			param.Page,
			param.Limit,
			int(attributes[0].TotalCount),
			int(attributes[0].CurrentCount),
		),
	}, nil
}

func (r *PostgresAttribute) Create(
	ctx context.Context,
	param service.CreateAttributeParam,
) (*domain.Attribute, error) {
	attribute, err := r.db.CreateAttribute(ctx,
		postgres.CreateAttributeParams{
			Code: param.Code,
			Name: param.Name,
		},
	)
	if err != nil {
		return nil, err
	}

	return mapPostgresAttributeToDomain(attribute), nil
}

func (r *PostgresAttribute) Update(
	ctx context.Context,
	param service.UpdateAttributeParam,
) (*domain.Attribute, error) {
	attribute, err := r.db.UpdateAttribute(ctx,
		postgres.UpdateAttributeParams{
			ID: int32(param.AttributeID),
			Name: pgtype.Text{
				String: param.Name,
				Valid:  true,
			},
			Code: pgtype.Text{
				Valid: false,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return mapPostgresAttributeToDomain(attribute), nil
}

func (r *PostgresAttribute) CreateValue(
	ctx context.Context,
	param service.CreateAttributeValueParam,
) (*domain.AttributeValue, error) {
	attributeValue, err := r.db.CreateAttributeValue(ctx,
		postgres.CreateAttributeValueParams{
			AttributeID: int32(param.AttributeID),
			Value:       param.Value,
		},
	)
	if err != nil {
		return nil, err
	}

	return mapPostgresAttributeValueToDomain(attributeValue), nil
}

func (r *PostgresAttribute) UpdateValues(
	ctx context.Context,
	param []service.UpdateAttributeValueParam,
) error {
	if len(param) == 0 {
		return nil
	}

	ids := domainSliceToPgtype(param, func(p service.UpdateAttributeValueParam) int32 {
		return int32(p.AttributeValueIds)
	})
	values := domainSliceToPgtype(param, func(p service.UpdateAttributeValueParam) string {
		return p.Values
	})

	_, err := r.db.UpdateAttributeValues(ctx,
		postgres.UpdateAttributeValuesParams{
			IDs:    ids,
			Values: values,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// Helper functions

func mapPostgresAttributeToDomain(attr postgres.Attribute) *domain.Attribute {
	return &domain.Attribute{
		ID:              int(attr.ID),
		Code:            attr.Code,
		Name:            attr.Name,
		AttributeValues: []domain.AttributeValue{},
		DeletedAt: mapPgtypeToDomainPtr(
			attr.DeletedAt.Time,
			attr.DeletedAt.Valid,
		),
	}
}

func mapListAttributeRowToDomain(attr postgres.ListAttributesRow) domain.Attribute {
	return domain.Attribute{
		ID:   int(attr.ID),
		Code: attr.Code,
		Name: attr.Name,
		DeletedAt: mapPgtypeToDomainPtr(
			attr.DeletedAt.Time,
			attr.DeletedAt.Valid,
		),
	}
}

func mapPostgresAttributeValueToDomain(val postgres.AttributeValue) *domain.AttributeValue {
	return &domain.AttributeValue{
		ID:    int(val.ID),
		Value: val.Value,
	}
}

func buildPaginationMeta(page, limit, totalCount, currentCount int) domain.PaginationMeta {
	return domain.PaginationMeta{
		TotalItems:   totalCount,
		CurrentPage:  page,
		ItemsPerPage: limit,
		PageItems:    currentCount,
	}
}
