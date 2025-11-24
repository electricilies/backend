package repositorypostgres

import (
	"context"
	"time"

	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Attribute struct {
	queries *sqlc.Queries
	conn    *pgxpool.Pool
}

var _ domain.AttributeRepository = (*Attribute)(nil)

func ProvideAttribute(q *sqlc.Queries, conn *pgxpool.Pool) *Attribute {
	return &Attribute{
		queries: q,
		conn:    conn,
	}
}

func (r *Attribute) Count(
	ctx context.Context,
	ids *[]uuid.UUID,
	deleted domain.DeletedParam,
) (*int, error) {
	count, err := r.queries.CountAttributes(ctx, sqlc.CountAttributesParams{
		IDs:     ptr.Deref(ids, []uuid.UUID{}),
		Deleted: string(deleted),
	})
	return ptr.To(int(count)), err
}

func (r *Attribute) List(
	ctx context.Context,
	ids *[]uuid.UUID,
	search *string,
	deleted domain.DeletedParam,
	limit int,
	offset int,
) (*[]domain.Attribute, error) {
	attributes, err := r.queries.ListAttributes(ctx, sqlc.ListAttributesParams{
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
	attributeValues, err := r.queries.ListAttributeValues(ctx, sqlc.ListAttributeValuesParams{
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
			Values: buildAttributeValues(
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

func (r *Attribute) ListValues(
	ctx context.Context,
	attributeID uuid.UUID,
	attributeValueIDs *[]uuid.UUID,
	search *string,
	deleted domain.DeletedParam,
	limit int,
	offset int,
) (*[]domain.AttributeValue, error) {
	attributeValues, err := r.queries.ListAttributeValues(
		ctx,
		sqlc.ListAttributeValuesParams{
			IDs: ptr.Deref(attributeValueIDs, []uuid.UUID{}),
			AttributeID: pgtype.UUID{
				Bytes: attributeID,
				Valid: true,
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

func (r *Attribute) CountValues(
	ctx context.Context,
	attributeID uuid.UUID,
	attributeValueIDs *[]uuid.UUID,
) (*int, error) {
	count, err := r.queries.CountAttributeValues(ctx, sqlc.CountAttributeValuesParams{
		IDs:     ptr.Deref(attributeValueIDs, []uuid.UUID{}),
		Deleted: string(domain.DeletedExcludeParam),
		AttributeID: pgtype.UUID{
			Bytes: attributeID,
			Valid: true,
		},
	})
	return ptr.To(int(count)), err
}

func (r *Attribute) Get(ctx context.Context, id uuid.UUID) (*domain.Attribute, error) {
	attribute, err := r.queries.GetAttribute(ctx, sqlc.GetAttributeParams{
		ID: id,
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}

	attributeValues, err := r.ListValues(ctx, id, nil, nil, domain.DeletedExcludeParam, 0, 0)
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

func (r *Attribute) Save(ctx context.Context, attribute domain.Attribute) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	qtx := r.queries.WithTx(tx)
	err = qtx.UpsertAttribute(ctx, sqlc.UpsertAttributeParams{
		ID:   attribute.ID,
		Code: attribute.Code,
		Name: attribute.Name,
		DeletedAt: pgtype.Timestamptz{
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
	_, err = qtx.InsertTempTableAttributeValues(ctx, buildInsertTempTableAttributeValuesParams(attribute))
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

func buildAttributeValues(attributeID uuid.UUID, attributeValues []sqlc.AttributeValue) []domain.AttributeValue {
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
	return values
}

func buildInsertTempTableAttributeValuesParams(attribute domain.Attribute) []sqlc.InsertTempTableAttributeValuesParams {
	params := make([]sqlc.InsertTempTableAttributeValuesParams, 0, len(attribute.Values))
	for _, val := range attribute.Values {
		params = append(params, sqlc.InsertTempTableAttributeValuesParams{
			ID:          val.ID,
			Value:       val.Value,
			AttributeID: attribute.ID,
		})
	}
	return params
}
