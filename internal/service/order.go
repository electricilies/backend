package service

import (
	"backend/internal/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

type Order struct {
	validate *validator.Validate
}

func ProvideOrder(
	validate *validator.Validate,
) *Order {
	return &Order{
		validate: validate,
	}
}

var _ domain.OrderService = &Order{}

func (o *Order) Create(
	address string,
	provider domain.OrderProvider,
	isPaid bool,
	totalAmount int64,
) (*domain.Order, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	order := &domain.Order{
		ID:          id,
		Address:     address,
		Provider:    provider,
		Status:      domain.OrderStatusPending,
		IsPaid:      isPaid,
		TotalAmount: totalAmount,
	}
	if err := o.validate.Struct(order); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return order, nil
}

func (o *Order) CreateItem(
	productVariant domain.ProductVariant,
	quantity int,
	price int64,
) (*domain.OrderItem, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	orderItem := &domain.OrderItem{
		ID:             id,
		ProductVariant: &productVariant,
		Quantity:       quantity,
		Price:          price,
	}
	if err := o.validate.Struct(orderItem); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return orderItem, nil
}
