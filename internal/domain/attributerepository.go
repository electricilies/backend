package domain

import (
	"context"

	"github.com/google/uuid"
)

type AttributeRepository interface {
	Count(
		ctx context.Context,
		params AttributeRepositoryCountParam,
	) (*int, error)

	List(
		ctx context.Context,
		params AttributeRepositoryListParam,
	) (*[]Attribute, error)

	ListValues(
		ctx context.Context,
		params AttributeRepositoryListValuesParam,
	) (*[]AttributeValue, error)

	CountValues(
		ctx context.Context,
		params AttributeRepositoryCountValuesParam,
	) (*int, error)

	Get(
		ctx context.Context,
		params AttributeRepositoryGetParam,
	) (*Attribute, error)

	Save(
		ctx context.Context,
		params AttributeRepositorySaveParam,
	) error
}

type AttributeRepositoryCountParam struct {
	IDs     []uuid.UUID
	Deleted DeletedParam
}

type AttributeRepositoryListParam struct {
	IDs               []uuid.UUID
	AttributeValueIDs []uuid.UUID
	Search            string
	Deleted           DeletedParam
	Limit             int
	Offset            int
}

type AttributeRepositoryListValuesParam struct {
	AttributeID       uuid.UUID
	AttributeValueIDs []uuid.UUID
	Search            string
	Deleted           DeletedParam
	Limit             int
	Offset            int
}

type AttributeRepositoryCountValuesParam struct {
	AttributeID       uuid.UUID
	AttributeValueIDs []uuid.UUID
}

type AttributeRepositoryGetParam struct {
	ID uuid.UUID
}

type AttributeRepositorySaveParam struct {
	Attribute Attribute
}
