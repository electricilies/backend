package domain

import (
	"context"

	"github.com/google/uuid"
)

type AttributeRepository interface {
	Count(
		ctx context.Context,
		ids *[]uuid.UUID,
		deleted DeletedParam,
	) (*int, error)

	List(
		ctx context.Context,
		ids *[]uuid.UUID,
		search *string,
		deleted DeletedParam,
		limit int,
		offset int,
	) (*[]Attribute, error)

	Get(
		ctx context.Context,
		id uuid.UUID,
	) (*Attribute, error)

	Save(
		ctx context.Context,
		attribute *Attribute,
	) error

	Remove(
		ctx context.Context,
		id uuid.UUID,
	) error
}
