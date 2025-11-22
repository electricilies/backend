package repository

import (
	"context"
	"time"

	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostgresAttribute struct {
	querier postgres.Querier
}

var _ domain.AttributeRepository = (*PostgresAttribute)(nil)

func ProvidePostgresAttribute(q postgres.Querier) *PostgresAttribute {
	return &PostgresAttribute{querier: q}
}

func (r *PostgresAttribute) Count(
	ctx context.Context,
	ids *[]uuid.UUID,
	deleted domain.DeletedParam,
) (*int, error) {
	count, err := r.querier.CountAttributes(ctx, postgres.CountAttributesParams{
		IDs:     ptr.Deref(ids, []uuid.UUID{}),
		Deleted: string(deleted),
	})
	return ptr.To(int(count)), err
}

func (r *PostgresAttribute) List(
	ctx context.Context,
	ids *[]uuid.UUID,
	search *string,
	deleted domain.DeletedParam,
	limit int,
	offset int,
) (*[]domain.Attribute, error) {
	attributes, err := r.querier.ListAttributes(ctx, postgres.ListAttributesParams{
		IDs:     ptr.Deref(ids, []uuid.UUID{}),
		Search:  search,
		Deleted: string(deleted),
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	result := make([]domain.Attribute, 0, len(attributes))
	for i, attr := range attributes {
		result[i] = domain.Attribute{
			ID:        attr.ID,
			Code:      attr.Code,
			Name:      attr.Name,
			DeletedAt: &attr.DeletedAt.Time,
		}
	}
	return &result, err
}

func (r *PostgresAttribute) ListValues(
	ctx context.Context,
	attributeID *uuid.UUID,
	attributeValueIDs *[]uuid.UUID,
	search *string,
	deleted domain.DeletedParam,
	limit int,
	offset int,
) (*[]domain.AttributeValue, error) {
	attributeValues, err := r.querier.ListAttributeValues(
		ctx,
		postgres.ListAttributeValuesParams{
			IDs: ptr.Deref(attributeValueIDs, []uuid.UUID{}),
			AttributeID: pgtype.UUID{
				Bytes: ptr.Deref(attributeID, uuid.UUID{}),
				Valid: attributeID != nil,
			},
			Search:  search,
			Deleted: string(deleted),
			Limit:   int32(limit),
			Offset:  int32(offset),
		},
	)
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	result := make([]domain.AttributeValue, 0, len(attributeValues))
	for i, attrVal := range attributeValues {
		result[i] = domain.AttributeValue{
			ID:        attrVal.ID,
			Value:     attrVal.Value,
			DeletedAt: &attrVal.DeletedAt.Time,
		}
	}
	return &result, nil
}

func (r *PostgresAttribute) CountValues(ctx context.Context, attributeID *uuid.UUID, attributeValueIDs *[]uuid.UUID) (*int, error) {
	count, err := r.querier.CountAttributeValues(ctx, postgres.CountAttributeValuesParams{
		IDs: ptr.Deref(attributeValueIDs, []uuid.UUID{}),
		AttributeID: pgtype.UUID{
			Bytes: ptr.Deref(attributeID, uuid.UUID{}),
			Valid: attributeID != nil,
		},
	})
	return ptr.To(int(count)), err
}

func (r *PostgresAttribute) Get(ctx context.Context, id uuid.UUID) (*domain.Attribute, error) {
	attribute, err := r.querier.GetAttribute(ctx, postgres.GetAttributeParams{
		ID: id,
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}

	attributeValues, err := r.ListValues(ctx, &id, nil, nil, domain.DeletedExcludeParam, 0, 0)
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}

	result := domain.Attribute{
		ID:        attribute.ID,
		Code:      attribute.Code,
		Name:      attribute.Name,
		Values:    attributeValues,
		DeletedAt: ptr.To(attribute.DeletedAt.Time),
	}
	return &result, nil
}

func (r *PostgresAttribute) Save(ctx context.Context, attribute domain.Attribute) error {
	err := r.querier.CreateTempTableAttributeValues(ctx)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	_, err = r.querier.InsertTempTableAttributeValues(ctx, func() []postgres.InsertTempTableAttributeValuesParams {
		params := make([]postgres.InsertTempTableAttributeValuesParams, 0, len(*attribute.Values))
		for _, val := range *attribute.Values {
			params = append(params, postgres.InsertTempTableAttributeValuesParams{
				ID:    val.ID,
				Value: val.Value,
			})
		}
		return params
	}())
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	err = r.querier.UpsertAttribute(ctx, postgres.UpsertAttributeParams{
		ID:   attribute.ID,
		Code: attribute.Code,
		Name: attribute.Name,
		DeletedAt: pgtype.Timestamp{
			Time:  ptr.Deref(attribute.DeletedAt, time.Time{}),
			Valid: attribute.DeletedAt != nil,
		},
	})
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	return nil
}
