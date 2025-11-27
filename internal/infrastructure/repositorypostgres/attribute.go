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
	params domain.AttributeRepositoryCountParam,
) (*int, error) {
	count, err := r.queries.CountAttributes(ctx, sqlc.CountAttributesParams{
		IDs:     params.IDs,
		Deleted: string(params.Deleted),
	})
	return ptr.To(int(count)), err
}

func (r *Attribute) List(
	ctx context.Context,
	params domain.AttributeRepositoryListParam,
) (*[]domain.Attribute, error) {
	attributes, err := r.queries.ListAttributes(ctx, sqlc.ListAttributesParams{
		IDs:               params.IDs,
		AttributeValueIDs: params.AttributeValueIDs,
		Search:            params.Search,
		Deleted:           string(params.Deleted),
		Limit:             int32(params.Limit),
		Offset:            int32(params.Offset),
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
		Deleted:      string(params.Deleted),
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
	params domain.AttributeRepositoryListValuesParam,
) (*[]domain.AttributeValue, error) {
	attributeValues, err := r.queries.ListAttributeValues(
		ctx,
		sqlc.ListAttributeValuesParams{
			IDs:         params.AttributeValueIDs,
			AttributeID: params.AttributeID,
			Search:      params.Search,
			Deleted:     string(params.Deleted),
			Limit:       int32(params.Limit),
			Offset:      int32(params.Offset),
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
	params domain.AttributeRepositoryCountValuesParam,
) (*int, error) {
	count, err := r.queries.CountAttributeValues(ctx, sqlc.CountAttributeValuesParams{
		IDs:         params.AttributeValueIDs,
		Deleted:     string(domain.DeletedExcludeParam),
		AttributeID: params.AttributeID,
	})
	return ptr.To(int(count)), err
}

func (r *Attribute) Get(ctx context.Context, params domain.AttributeRepositoryGetParam) (*domain.Attribute, error) {
	attribute, err := r.queries.GetAttribute(ctx, sqlc.GetAttributeParams{
		ID: params.ID,
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}

	attributeValues, err := r.ListValues(ctx, domain.AttributeRepositoryListValuesParam{
		AttributeID: params.ID,
		Deleted:     domain.DeletedExcludeParam,
	})
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

func (r *Attribute) Save(ctx context.Context, params domain.AttributeRepositorySaveParam) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	qtx := r.queries.WithTx(tx)
	err = qtx.UpsertAttribute(ctx, sqlc.UpsertAttributeParams{
		ID:   params.Attribute.ID,
		Code: params.Attribute.Code,
		Name: params.Attribute.Name,
		DeletedAt: pgtype.Timestamptz{
			Time:  ptr.Deref(params.Attribute.DeletedAt, time.Time{}),
			Valid: params.Attribute.DeletedAt != nil,
		},
	})
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	err = qtx.CreateTempTableAttributeValues(ctx)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	_, err = qtx.InsertTempTableAttributeValues(ctx, buildInsertTempTableAttributeValuesParams(params.Attribute))
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
