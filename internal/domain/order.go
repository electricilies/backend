package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID     `validate:"required"`
	Address     string        `validate:"required"`
	Provider    OrderProvider `validate:"required,oneof=COD VNPAY MOMO ZALOPAY"`
	Status      OrderStatus   `validate:"required,oneof=Pending Processing Shipped Delivered Cancelled"`
	IsPaid      bool          `validate:"required"`
	CreatedAt   time.Time     `validate:"required"`
	UpdatedAt   time.Time     `validate:"required,gtefield=CreatedAt"`
	Items       []OrderItem   `validate:"omitempty,dive"`
	TotalAmount int64         `validate:"required"`
	UserID      uuid.UUID     `validate:"required"`
}

type OrderItem struct {
	ID               uuid.UUID `validate:"required"`
	ProductID        uuid.UUID `validate:"required"`
	ProductVariantID uuid.UUID `validate:"required"`
	Quantity         int       `validate:"required,gt=0"`
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
	now := time.Now()
	return &Order{
		ID:        id,
		UserID:    userID,
		Address:   address,
		Provider:  provider,
		Status:    OrderStatusPending,
		IsPaid:    false,
		CreatedAt: now,
		UpdatedAt: now,
		Items:     items,
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

func (o *Order) AddItems(items ...OrderItem) {
	o.Items = append(o.Items, items...)
}
