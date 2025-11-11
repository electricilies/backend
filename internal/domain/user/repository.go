package user

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, id string) (*Model, error)
	List(ctx context.Context) ([]*Model, error)
	Create(ctx context.Context, user *Model) (*Model, error)
	Update(ctx context.Context, user *Model) error
	Delete(ctx context.Context, id string) error
}
