package service

import (
	"context"

	"backend/internal/domain"
)

type CreatePaymentParam struct {
	Provider string `json:"paymentProvider" binding:"required"`
	OrderID  int    `json:"orderId" binding:"required"`
}

type UpdatePaymentParam struct {
	Status string `json:"status" binding:"required"`
}

type ListPaymentParam struct {
	PaginationParam
}

type GetPaymentParam struct {
	PaymentID int `json:"paymentId" binding:"required"`
}

type Payment interface {
	Create(context.Context, CreatePaymentParam) (*domain.Payment, error)
	Update(context.Context, UpdatePaymentParam) (*domain.Payment, error)
	List(context.Context, ListPaymentParam) (*domain.Pagination[domain.Payment], error)
	Get(context.Context, int) (*domain.Payment, error)
}

type PaymentImpl struct{}

func ProvidePayment() *PaymentImpl {
	return &PaymentImpl{}
}

var _ Payment = &PaymentImpl{}

func (s *PaymentImpl) Create(ctx context.Context, param CreatePaymentParam) (*domain.Payment, error) {
	return nil, nil
}

func (s *PaymentImpl) Update(ctx context.Context, param UpdatePaymentParam) (*domain.Payment, error) {
	return nil, nil
}

func (s *PaymentImpl) List(ctx context.Context, param ListPaymentParam) (*domain.Pagination[domain.Payment], error) {
	return nil, nil
}

func (s *PaymentImpl) Get(ctx context.Context, paymentID int) (*domain.Payment, error) {
	return nil, nil
}
