package payment

import "context"

type Repository interface {
	List(context.Context) ([]*Model, error)
	Create(context.Context, *Model) (*Model, error)
	Get(context.Context, int) (*Model, error)
	Update(context.Context, *Model, int) (*Model, error)
}
