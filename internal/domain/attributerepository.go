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

	ListValues(
		ctx context.Context,
		attributeID *uuid.UUID,
		attributeValueIDs *[]uuid.UUID,
		search *string,
		deleted DeletedParam,
		limit int,
		offset int,
	) (*[]AttributeValue, error)

	CountValues(
		ctx context.Context,
		attributeID *uuid.UUID,
		attributeValueIDs *[]uuid.UUID,
	) (*int, error)

	Get(
		ctx context.Context,
		id uuid.UUID,
	) (*Attribute, error)

	Save(
		ctx context.Context,
		attribute Attribute,
	) error
}
