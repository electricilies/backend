package repository

import (
	"context"
	"log"
	"time"

	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresAttribute struct {
	queries *postgres.Queries
	conn    *pgxpool.Pool
}

var _ domain.AttributeRepository = (*PostgresAttribute)(nil)

func ProvidePostgresAttribute(q *postgres.Queries, conn *pgxpool.Pool) *PostgresAttribute {
	return &PostgresAttribute{
		queries: q,
		conn:    conn,
	}
}

func (r *PostgresAttribute) Count(
	ctx context.Context,
	ids *[]uuid.UUID,
	deleted domain.DeletedParam,
) (*int, error) {
	count, err := r.queries.CountAttributes(ctx, postgres.CountAttributesParams{
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
	attributes, err := r.queries.ListAttributes(ctx, postgres.ListAttributesParams{
		IDs:     ptr.Deref(ids, []uuid.UUID{}),
		Search:  search,
		Deleted: string(deleted),
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	attributeIDs := make([]uuid.UUID, 0, len(attributes))
	for _, attr := range attributes {
		attributeIDs = append(attributeIDs, attr.ID)
	}
	attributeValues, err := r.queries.ListAttributeValues(ctx, postgres.ListAttributeValuesParams{
		AttributeIDs: attributeIDs,
		Deleted:      string(deleted),
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	result := make([]domain.Attribute, 0, len(attributes))
	for _, attribute := range attributes {
		result = append(result, domain.Attribute{
			ID:   attribute.ID,
			Code: attribute.Code,
			Name: attribute.Name,
			Values: *buildAttributeValues(
				attribute.ID,
				attributeValues,
			),
			DeletedAt: fromPgValidToPtr(
				attribute.DeletedAt.Time,
				attribute.DeletedAt.Valid,
			),
		})
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
	attributeValues, err := r.queries.ListAttributeValues(
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
	for _, attrVal := range attributeValues {
		result = append(result, domain.AttributeValue{
			ID:    attrVal.ID,
			Value: attrVal.Value,
			DeletedAt: fromPgValidToPtr(
				attrVal.DeletedAt.Time,
				attrVal.DeletedAt.Valid,
			),
		})
	}
	return &result, nil
}

func (r *PostgresAttribute) CountValues(ctx context.Context, attributeID *uuid.UUID, attributeValueIDs *[]uuid.UUID) (*int, error) {
	count, err := r.queries.CountAttributeValues(ctx, postgres.CountAttributeValuesParams{
		IDs:     ptr.Deref(attributeValueIDs, []uuid.UUID{}),
		Deleted: string(domain.DeletedExcludeParam),
		AttributeID: pgtype.UUID{
			Bytes: ptr.Deref(attributeID, uuid.UUID{}),
			Valid: attributeID != nil,
		},
	})
	return ptr.To(int(count)), err
}

func (r *PostgresAttribute) Get(ctx context.Context, id uuid.UUID) (*domain.Attribute, error) {
	attribute, err := r.queries.GetAttribute(ctx, postgres.GetAttributeParams{
		ID: id,
	})
	log.Println("Fetched attribute:", attribute, "Id", id)
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}

	attributeValues, err := r.ListValues(ctx, &id, nil, nil, domain.DeletedExcludeParam, 0, 0)
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}

	result := domain.Attribute{
		ID:     attribute.ID,
		Code:   attribute.Code,
		Name:   attribute.Name,
		Values: *attributeValues,
		DeletedAt: fromPgValidToPtr(
			attribute.DeletedAt.Time,
			attribute.DeletedAt.Valid,
		),
	}
	return &result, nil
}

func (r *PostgresAttribute) Save(ctx context.Context, attribute domain.Attribute) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	qtx := r.queries.WithTx(tx)
	err = qtx.UpsertAttribute(ctx, postgres.UpsertAttributeParams{
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
	err = qtx.CreateTempTableAttributeValues(ctx)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	_, err = qtx.InsertTempTableAttributeValues(ctx, buildInsertTempTableAttributeValuesParams(ptr.To(attribute.Values)))
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	err = qtx.MergeAttributeValuesFromTemp(ctx)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	return nil
}

func buildAttributeValues(attributeID uuid.UUID, attributeValues []postgres.AttributeValue) *[]domain.AttributeValue {
	values := make([]domain.AttributeValue, 0)
	for _, val := range attributeValues {
		if val.AttributeID == attributeID {
			values = append(values, domain.AttributeValue{
				ID:    val.ID,
				Value: val.Value,
				DeletedAt: fromPgValidToPtr(
					val.DeletedAt.Time,
					val.DeletedAt.Valid,
				),
			})
		}
	}
	return &values
}

func buildInsertTempTableAttributeValuesParams(values *[]domain.AttributeValue) []postgres.InsertTempTableAttributeValuesParams {
	params := make([]postgres.InsertTempTableAttributeValuesParams, 0, len(*values))
	for _, val := range *values {
		params = append(params, postgres.InsertTempTableAttributeValuesParams{
			ID:    val.ID,
			Value: val.Value,
		})
	}
	return params
}
