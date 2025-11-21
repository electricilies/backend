package domain

import (
	"context"
)

type AttributeRepository interface {
	Count(
		ctx context.Context,
		ids *[]int,
		deleted string,
		limit int,
		offset int,
	) (*int, error)

	List(
		ctx context.Context,
		ids *[]int,
		search *string,
		deleted string,
		limit int,
		offset int,
	) (*[]Attribute, error)

	Get(
		ctx context.Context,
		id int,
	) (*Attribute, error)

	Save(
		ctx context.Context,
		attribute *Attribute,
	) (*Attribute, error)

	Remove(
		ctx context.Context,
		id int,
	) error
}
