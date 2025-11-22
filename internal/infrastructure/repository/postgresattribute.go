package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresAttribute struct {
	querier postgres.Querier
}

var _ domain.AttributeRepository = (*PostgresAttribute)(nil)

func ProvidePostgresAttribute(q postgres.Querier) *PostgresAttribute {
	return &PostgresAttribute{querier: q}
}

func (r *PostgresAttribute) Count(ctx context.Context, ids *[]uuid.UUID, deleted domain.DeletedParam) (*int, error) {
	panic("implement me")
}

func (r *PostgresAttribute) List(ctx context.Context, ids *[]uuid.UUID, search *string, deleted domain.DeletedParam, limit int, offset int) (*[]domain.Attribute, error) {
	panic("implement me")
}

func (r *PostgresAttribute) ListValues(ctx context.Context, attributeID *uuid.UUID, attributeValueIDs *[]uuid.UUID, search *string, limit int, offset int) (*[]domain.AttributeValue, error) {
	panic("implement me")
}

func (r *PostgresAttribute) CountValues(ctx context.Context, attributeID *uuid.UUID, attributeValueIDs *[]uuid.UUID) (*int, error) {
	panic("implement me")
}

func (r *PostgresAttribute) Get(ctx context.Context, id uuid.UUID) (*domain.Attribute, error) {
	panic("implement me")
}

func (r *PostgresAttribute) Save(ctx context.Context, attribute domain.Attribute) error {
	panic("implement me")
}
