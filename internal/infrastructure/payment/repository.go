package payment

import (
	"context"

	"backend/internal/domain/payment"
)

type RepositoryImpl struct{}

func NewRepository() payment.Repository {
	return &RepositoryImpl{}
}

func ProvideRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) List(ctx context.Context) ([]*payment.Model, error) {
	return []*payment.Model{}, nil
}

func (r *RepositoryImpl) Create(ctx context.Context, model *payment.Model) (*payment.Model, error) {
	return model, nil
}

func (r *RepositoryImpl) Get(ctx context.Context, id int) (*payment.Model, error) {
	return &payment.Model{}, nil
}

func (r *RepositoryImpl) Update(ctx context.Context, model *payment.Model, id int) (*payment.Model, error) {
	return model, nil
}
