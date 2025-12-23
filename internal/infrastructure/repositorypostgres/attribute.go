package repositorypostgres

import (
	"context"

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
		IDs:               params.IDs,
		Search:            params.Search,
		AttributeValueIDs: params.AttributeValueIDs,
		Deleted:           string(params.Deleted),
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
		return nil, toDomainError(err)
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
		return nil, toDomainError(err)
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
			DeletedAt: attribute.DeletedAt.Time,
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
			IDs:          params.IDs,
			AttributeID:  params.AttributeID,
			AttributeIDs: params.AttributeIDs,
			Search:       params.Search,
			Deleted:      string(params.Deleted),
			Limit:        int32(params.Limit),
			Offset:       int32(params.Offset),
		},
	)
	if err != nil {
		return nil, toDomainError(err)
	}
	result := make([]domain.AttributeValue, 0, len(attributeValues))
	for _, attrVal := range attributeValues {
		result = append(result, domain.AttributeValue{
			ID:        attrVal.ID,
			Value:     attrVal.Value,
			DeletedAt: attrVal.DeletedAt.Time,
		})
	}
	return &result, nil
}

func (r *Attribute) CountValues(
	ctx context.Context,
	params domain.AttributeRepositoryCountValuesParam,
) (*int, error) {
	count, err := r.queries.CountAttributeValues(ctx, sqlc.CountAttributeValuesParams{
		IDs:          params.IDs,
		AttributeID:  params.AttributeID,
		AttributeIDs: params.AttributeIDs,
		Search:       params.Search,
		Deleted:      string(params.Deleted),
	})
	return ptr.To(int(count)), err
}

func (r *Attribute) Get(ctx context.Context, params domain.AttributeRepositoryGetParam) (*domain.Attribute, error) {
	attribute, err := r.queries.GetAttribute(ctx, sqlc.GetAttributeParams{
		ID: params.ID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	attributeValues, err := r.ListValues(ctx, domain.AttributeRepositoryListValuesParam{
		AttributeID: params.ID,
		Deleted:     domain.DeletedExcludeParam,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	result := domain.Attribute{
		ID:        attribute.ID,
		Code:      attribute.Code,
		Name:      attribute.Name,
		Values:    *attributeValues,
		DeletedAt: attribute.DeletedAt.Time,
	}
	return &result, nil
}

func (r *Attribute) Save(ctx context.Context, params domain.AttributeRepositorySaveParam) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return toDomainError(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	qtx := r.queries.WithTx(tx)
	err = qtx.UpsertAttribute(ctx, sqlc.UpsertAttributeParams{
		ID:   params.Attribute.ID,
		Code: params.Attribute.Code,
		Name: params.Attribute.Name,
		DeletedAt: pgtype.Timestamptz{
			Time:  params.Attribute.DeletedAt,
			Valid: !params.Attribute.DeletedAt.IsZero(),
		},
	})
	if err != nil {
		return toDomainError(err)
	}
	err = qtx.CreateTempTableAttributeValues(ctx)
	if err != nil {
		return toDomainError(err)
	}
	_, err = qtx.InsertTempTableAttributeValues(ctx, buildInsertTempTableAttributeValuesParams(params.Attribute))
	if err != nil {
		return toDomainError(err)
	}
	err = qtx.MergeAttributeValuesFromTemp(ctx)
	if err != nil {
		return toDomainError(err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return toDomainError(err)
	}
	return nil
}

func buildAttributeValues(attributeID uuid.UUID, attributeValues []sqlc.AttributeValue) []domain.AttributeValue {
	values := make([]domain.AttributeValue, 0)
	for _, val := range attributeValues {
		if val.AttributeID == attributeID {
			values = append(values, domain.AttributeValue{
				ID:        val.ID,
				Value:     val.Value,
				DeletedAt: val.DeletedAt.Time,
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
