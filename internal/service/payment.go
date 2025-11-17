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
	Limit  int
	Offset int
}

type GetPaymentParam struct {
	PaymentID int `json:"paymentId" binding:"required"`
}

type Payment interface {
	CreatePayment(context.Context, CreatePaymentParam) (*domain.Payment, error)
	UpdatePayment(context.Context, UpdatePaymentParam) (*domain.Payment, error)
	ListPayments(context.Context, ListPaymentParam) (*domain.DataPagination, error)
	GetPayment(context.Context, int) (*domain.Payment, error)
}

type PaymentImpl struct{}

func ProvidePayment() *PaymentImpl {
	return &PaymentImpl{}
}

var _ Payment = &PaymentImpl{}

func (s *PaymentImpl) CreatePayment(ctx context.Context, param CreatePaymentParam) (*domain.Payment, error) {
	return nil, nil
}

func (s *PaymentImpl) UpdatePayment(ctx context.Context, param UpdatePaymentParam) (*domain.Payment, error) {
	return nil, nil
}

func (s *PaymentImpl) ListPayments(ctx context.Context, param ListPaymentParam) (*domain.DataPagination, error) {
	return nil, nil
}

func (s *PaymentImpl) GetPayment(ctx context.Context, paymentID int) (*domain.Payment, error) {
	return nil, nil
}

