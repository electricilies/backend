package application

import (
	"context"

	"backend/internal/domain"
)

type OrderImpl struct {
	orderRepo    domain.OrderRepository
	orderService domain.OrderService
}

func ProvideOrder(orderRepo domain.OrderRepository, orderService domain.OrderService) *OrderImpl {
	return &OrderImpl{
		orderRepo:    orderRepo,
		orderService: orderService,
	}
}

var _ Order = &OrderImpl{}

func (o *OrderImpl) Create(ctx context.Context, param CreateOrderParam) (*domain.Order, error) {
	// Convert CreateOrderItemData to OrderItems
	items := make([]domain.OrderItem, 0, len(param.Data.Items))
	for _, itemData := range param.Data.Items {
		item, err := o.orderService.CreateItem(itemData.ProductVariantID, itemData.Quantity, itemData.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}
	
	order, err := o.orderService.Create(
		param.UserID,
		param.Data.Address,
		param.Data.Provider,
		param.Data.TotalAmount,
		items,
	)
	if err != nil {
		return nil, err
	}
	
	err = o.orderRepo.Save(ctx, order)
	if err != nil {
		return nil, err
	}
	
	return order, nil
}

func (o *OrderImpl) Update(ctx context.Context, param UpdateOrderParam) (*domain.Order, error) {
	order, err := o.orderRepo.Get(ctx, param.OrderID)
	if err != nil {
		return nil, err
	}
	
	err = o.orderService.Update(order, param.Data.Status, param.Data.Address)
	if err != nil {
		return nil, err
	}
	
	err = o.orderRepo.Save(ctx, order)
	if err != nil {
		return nil, err
	}
	
	return order, nil
}

func (o *OrderImpl) Get(ctx context.Context, param GetOrderParam) (*domain.Order, error) {
	order, err := o.orderRepo.Get(ctx, param.OrderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *OrderImpl) Delete(ctx context.Context, param DeleteOrderParam) error {
	return o.orderRepo.Remove(ctx, param.OrderID)
}

func (o *OrderImpl) List(ctx context.Context, param ListOrderParam) (*Pagination[domain.Order], error) {
	orders, err := o.orderRepo.List(
		ctx,
		param.IDs,
		param.UserIDs,
		param.StatusIDs,
		*param.Limit,
		*param.Page,
	)
	if err != nil {
		return nil, err
	}
	
	count, err := o.orderRepo.Count(
		ctx,
		param.IDs,
		param.UserIDs,
		param.StatusIDs,
	)
	if err != nil {
		return nil, err
	}
	
	pagination := newPagination(*orders, *count, *param.Page, *param.Limit)
	return pagination, nil
}
