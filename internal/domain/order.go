package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID     `validate:"required"`
	Address     string        `validate:"required"`
	Provider    OrderProvider `validate:"required"`
	Status      OrderStatus   `validate:"required"`
	IsPaid      bool          `validate:"required"`
	CreatedAt   time.Time     `validate:"required"`
	UpdatedAt   time.Time     `validate:"required,gtefield=CreatedAt"`
	Items       []OrderItem   `validate:"gt=0,orderTotalAmount,dive"`
	TotalAmount int64         `validate:"required"`
	UserID      uuid.UUID     `validate:"required"`
}

type OrderItem struct {
	ID               uuid.UUID `validate:"required"`
	ProductID        uuid.UUID `validate:"required"`
	ProductVariantID uuid.UUID `validate:"required"`
	Quantity         int       `validate:"required,gt=0,lt=100"`
	Price            int64     `validate:"required,gt=0"`
}

type OrderProvider string

const (
	PaymentProviderCOD     OrderProvider = "COD"
	PaymentProviderVNPAY   OrderProvider = "VNPAY"
	PaymentProviderMOMO    OrderProvider = "MOMO"
	PaymentProviderZALOPAY OrderProvider = "ZALOPAY"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "Pending"
	OrderStatusProcessing OrderStatus = "Processing"
	OrderStatusShipped    OrderStatus = "Shipped"
	OrderStatusDelivered  OrderStatus = "Delivered"
	OrderStatusCancelled  OrderStatus = "Cancelled"
)

func NewOrder(
	userID uuid.UUID,
	address string,
	provider OrderProvider,
	items []OrderItem,
) (*Order, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	totalAmount := int64(0)
	for _, item := range items {
		totalAmount += item.Price * int64(item.Quantity)
	}
	now := time.Now()
	return &Order{
		ID:          id,
		UserID:      userID,
		Address:     address,
		Provider:    provider,
		Status:      OrderStatusPending,
		IsPaid:      false,
		CreatedAt:   now,
		UpdatedAt:   now,
		Items:       items,
		TotalAmount: totalAmount,
	}, nil
}

func NewOrderItem(
	productID uuid.UUID,
	productVariantID uuid.UUID,
	quantity int,
	price int64,
) (*OrderItem, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &OrderItem{
		ID:               id,
		ProductID:        productID,
		ProductVariantID: productVariantID,
		Quantity:         quantity,
		Price:            price,
	}, nil
}

func (o *Order) Update(
	address string,
	status OrderStatus,
	isPaid bool,
) {
	updated := false
	if o.Address != address {
		o.Address = address
		updated = true
	}
	if o.Status != status {
		o.Status = status
		updated = true
	}
	if o.IsPaid != isPaid {
		o.IsPaid = isPaid
		updated = true
	}
	if updated {
		o.UpdatedAt = time.Now()
	}
}
