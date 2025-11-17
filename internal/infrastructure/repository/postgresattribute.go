package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/repository"
	"backend/internal/service"
)

type PostgresAttribute struct{}

func ProvidePostgresAttribute() *PostgresAttribute {
	return &PostgresAttribute{}
}

var _ repository.Attribute = &PostgresAttribute{}

func (r *PostgresAttribute) GetAttribute(
	ctx context.Context,
	query service.GetAttributeParam,
) (*domain.Attribute, error) {
	return nil, nil
}

func (r *PostgresAttribute) ListAttributes(
	ctx context.Context,
	query service.ListAttributesParam,
) (*[]domain.Attribute, error) {
	return nil, nil
}

func (r *PostgresAttribute) CreateAttribute(
	ctx context.Context,
	cmd service.CreateAttributeParam,
) (*domain.Attribute, error) {
	return nil, nil
}

func (r *PostgresAttribute) UpdateAttribute(
	ctx context.Context,
	cmd service.UpdateAttributeParam,
) (*domain.Attribute, error) {
	return nil, nil
}

func (r *PostgresAttribute) CreateAttributeValue(
	ctx context.Context,
	cmd service.CreateAttributeValueParam,
) (*domain.AttributeValue, error) {
	return nil, nil
}

func (r *PostgresAttribute) UpdateAttributeValues(
	ctx context.Context,
	cmd []service.UpdateAttributeValueParam,
) error {
	return nil
}
