package attribute

import (
	"backend/internal/domain/common"
	"context"
)

type AttributeRepository interface {
	Create(ctx context.Context, attribute Attribute) (*Attribute, error)
	List(ctx context.Context, param ListParam) (*[]Attribute, error)
	Get(context.Context, GetParam) (*Attribute, error)
	Update(ctx context.Context, attribute Attribute) (*Attribute, error)
	Delete(ctx context.Context, attribute Attribute) error
}

type ListParam struct {
	common.PaginationParam
	ProductID int
	Search    string
	Deleted   common.DeletedParam
}

type GetParam struct {
	ID string
}
