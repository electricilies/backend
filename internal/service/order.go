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
	userID uuid.UUID,
	address string,
	provider domain.OrderProvider,
	totalAmount int64,
	items []domain.OrderItem,
) (*domain.Order, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	order := &domain.Order{
		ID:          id,
		UserID:      userID,
		Address:     address,
		Provider:    provider,
		Status:      domain.OrderStatusPending,
		IsPaid:      false,
		TotalAmount: totalAmount,
		Items:       items,
	}
	if err := o.validate.Struct(order); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return order, nil
}

func (o *Order) CreateItem(
	productID uuid.UUID,
	productVariantID uuid.UUID,
	quantity int,
	price int64,
) (*domain.OrderItem, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	orderItem := &domain.OrderItem{
		ID:               id,
		ProductID:        productID,
		ProductVariantID: productVariantID,
		Quantity:         quantity,
		Price:            price,
	}
	if err := o.validate.Struct(orderItem); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return orderItem, nil
}

func (o *Order) Update(
	order *domain.Order,
	status *domain.OrderStatus,
	address *string,
) error {
	if status != nil {
		order.Status = *status
	}
	if address != nil {
		order.Address = *address
	}
	if err := o.validate.Struct(order); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
