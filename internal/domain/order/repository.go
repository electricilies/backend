package order

import "context"

type Repository interface {
	List(ctx context.Context) (*[]Model, error)
	Create(ctx context.Context, order *Model) (*Model, error)
	Update(ctx context.Context, order *Model, id int) (*Model, error)
	Delete(ctx context.Context, id int) error
}
