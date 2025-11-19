package serviceimpl

import (
	"context"

	"backend/internal/domain"
)

type Payment struct{}

func ProvidePayment() *Payment {
	return &Payment{}
}

var _ domain.PaymentService = &Payment{}

func (s *Payment) Create(ctx context.Context, param domain.CreatePaymentParam) (*domain.Payment, error) {
	panic("implement me")
}

func (s *Payment) Update(ctx context.Context, param domain.UpdatePaymentParam) (*domain.Payment, error) {
	panic("implement me")
}

func (s *Payment) Get(ctx context.Context, param domain.GetPaymentParam) (*domain.Payment, error) {
	panic("implement me")
}
