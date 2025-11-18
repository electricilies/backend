package service

import (
	"context"

	"backend/internal/domain"
)

type PaymentImpl struct{}

func ProvidePayment() *PaymentImpl {
	return &PaymentImpl{}
}

var _ Payment = &PaymentImpl{}

func (s *PaymentImpl) Create(ctx context.Context, param CreatePaymentParam) (*domain.Payment, error) {
	panic("implement me")
}

func (s *PaymentImpl) Update(ctx context.Context, param UpdatePaymentParam) (*domain.Payment, error) {
	panic("implement me")
}

func (s *PaymentImpl) List(ctx context.Context, param ListPaymentParam) (*Pagination[domain.Payment], error) {
	panic("implement me")
}

func (s *PaymentImpl) Get(ctx context.Context, param GetPaymentParam) (*domain.Payment, error) {
	panic("implement me")
}
