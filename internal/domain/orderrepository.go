package order

import "context"

type Repository interface {
	List(context.Context) (*[]Model, error)
	Create(context.Context, *Model) (*Model, error)
	Update(context.Context, *Model, int) (*Model, error)
	Delete(context.Context, int) error
}
