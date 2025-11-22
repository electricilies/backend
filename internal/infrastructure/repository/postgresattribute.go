package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresAttributes struct {
	querier postgres.Querier
}

func NewPostgresAttributes(q postgres.Querier) *PostgresAttributes {
	return &PostgresAttributes{querier: q}
}

func (r *PostgresAttributes) Count(ctx context.Context, ids *[]uuid.UUID, deleted domain.DeletedParam) (*int, error) {
	panic("implement me")
}

func (r *PostgresAttributes) List(ctx context.Context, ids *[]uuid.UUID, search *string, deleted domain.DeletedParam, limit int, offset int) (*[]domain.Attribute, error) {
	panic("implement me")
}

func (r *PostgresAttributes) ListValues(ctx context.Context, attributeID *uuid.UUID, attributeValueIDs *[]uuid.UUID, search *string, limit int, offset int) (*[]domain.AttributeValue, error) {
	panic("implement me")
}

func (r *PostgresAttributes) CountValues(ctx context.Context, attributeID *uuid.UUID, attributeValueIDs *[]uuid.UUID) (*int, error) {
	panic("implement me")
}

func (r *PostgresAttributes) Get(ctx context.Context, id uuid.UUID) (*domain.Attribute, error) {
	panic("implement me")
}

func (r *PostgresAttributes) Save(ctx context.Context, attribute domain.Attribute) error {
	panic("implement me")
}
