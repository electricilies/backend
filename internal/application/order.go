package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
)

type Order struct {
	orderRepo      domain.OrderRepository
	orderService   domain.OrderService
	productRepo    domain.ProductRepository
	productService domain.ProductService
}

func ProvideOrder(
	orderRepo domain.OrderRepository,
	orderService domain.OrderService,
	productRepo domain.ProductRepository,
	productService domain.ProductService,
) *Order {
	return &Order{
		orderRepo:      orderRepo,
		orderService:   orderService,
		productRepo:    productRepo,
		productService: productService,
	}
}

var _ http.OrderApplication = &Order{}

func (o *Order) Create(ctx context.Context, param http.CreateOrderRequestDto) (*http.OrderResponseDto, error) {
	productIDs := make([]uuid.UUID, 0, len(param.Data.Items))
	for _, item := range param.Data.Items {
		productIDs = append(productIDs, item.ProductID)
	}
	productVariantIDs := make([]uuid.UUID, 0, len(param.Data.Items))
	for _, item := range param.Data.Items {
		productVariantIDs = append(productVariantIDs, item.ProductVariantID)
	}
	products, err := o.productRepo.List(ctx, domain.ProductRepositoryListParam{
		IDs: productIDs,
	})
	if err != nil {
		return nil, err
	}
	productVariants, err := o.productService.FilterProductVariantsInProducts(
		*products,
		productVariantIDs,
	)
	if err != nil {
		return nil, err
	}
	productVariantIDProductVariantMap := make(map[uuid.UUID]domain.ProductVariant, len(*productVariants))
	for _, variant := range *productVariants {
		productVariantIDProductVariantMap[variant.ID] = variant
	}
	items := make([]domain.OrderItem, 0, len(param.Data.Items))
	for _, itemData := range param.Data.Items {
		variant, ok := productVariantIDProductVariantMap[itemData.ProductVariantID]
		if !ok {
			return nil, domain.ErrNotFound
		}
		item, err := domain.NewOrderItem(
			itemData.ProductID,
			itemData.ProductVariantID,
			itemData.Quantity,
			variant.Price,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}
	order, err := domain.NewOrder(
		param.Data.UserID,
		param.Data.Address,
		param.Data.Provider,
		items,
	)
	if err != nil {
		return nil, err
	}

	err = o.orderRepo.Save(ctx, domain.OrderRepositorySaveParam{
		Order: *order,
	})
	if err != nil {
		return nil, err
	}

	orderDto := http.ToOrderResponseDto(order)
	if err := o.enrichOrderItems(ctx, orderDto, order); err != nil {
		return nil, err
	}

	return orderDto, nil
}

func (o *Order) enrichOrderItems(ctx context.Context, orderDto *http.OrderResponseDto, order *domain.Order) error {
	if len(order.Items) == 0 {
		return nil
	}

	// Collect unique product IDs and variant IDs
	productIDsMap := make(map[uuid.UUID]struct{})
	variantIDs := make([]uuid.UUID, 0, len(order.Items))
	for _, item := range order.Items {
		productIDsMap[item.ProductID] = struct{}{}
		variantIDs = append(variantIDs, item.ProductVariantID)
	}

	productIDs := make([]uuid.UUID, 0, len(productIDsMap))
	for id := range productIDsMap {
		productIDs = append(productIDs, id)
	}

	// Fetch all products at once
	products, err := o.productRepo.List(
		ctx,
		domain.ProductRepositoryListParam{
			IDs:     productIDs,
			Deleted: domain.DeletedExcludeParam,
		},
	)
	if err != nil {
		return err
	}

	// Get filtered variants using product service
	productVariants, err := o.productService.FilterProductVariantsInProducts(*products, variantIDs)
	if err != nil {
		return err
	}

	// Create maps for quick lookup
	productMap := make(map[uuid.UUID]*domain.Product)
	for i := range *products {
		productMap[(*products)[i].ID] = &(*products)[i]
	}

	variantMap := make(map[uuid.UUID]*domain.ProductVariant)
	for i := range *productVariants {
		variantMap[(*productVariants)[i].ID] = &(*productVariants)[i]
	}

	// Build enriched order items
	enrichedItems := make([]http.OrderItemResponseDto, 0, len(order.Items))
	for _, item := range order.Items {
		product := productMap[item.ProductID]
		variant := variantMap[item.ProductVariantID]

		itemDto := http.ToOrderItemResponseDto(&item, product, variant)
		if itemDto != nil {
			enrichedItems = append(enrichedItems, *itemDto)
		}
	}

	orderDto.WithOrderItems(enrichedItems)
	return nil
}

func (o *Order) List(ctx context.Context, param http.ListOrderRequestDto) (*http.PaginationResponseDto[http.OrderResponseDto], error) {
	orders, err := o.orderRepo.List(
		ctx,
		domain.OrderRepositoryListParam{
			IDs:     param.IDs,
			Deleted: domain.DeletedExcludeParam,
			Limit:   param.Limit,
			Offset:  (param.Page - 1) * param.Limit,
		},
	)
	if err != nil {
		return nil, err
	}

	count, err := o.orderRepo.Count(
		ctx,
		domain.OrderRepositoryCountParam{
			IDs:     param.IDs,
			Deleted: domain.DeletedExcludeParam,
		},
	)
	if err != nil {
		return nil, err
	}

	orderDtos := make([]http.OrderResponseDto, 0, len(*orders))
	for i := range *orders {
		order := &(*orders)[i]
		orderDto := http.ToOrderResponseDto(order)
		if err := o.enrichOrderItems(ctx, orderDto, order); err != nil {
			return nil, err
		}
		orderDtos = append(orderDtos, *orderDto)
	}

	pagination := newPaginationResponseDto(
		orderDtos,
		*count,
		param.Page,
		param.Limit,
	)
	return pagination, nil
}

func (o *Order) Get(ctx context.Context, param http.GetOrderRequestDto) (*http.OrderResponseDto, error) {
	order, err := o.orderRepo.Get(ctx, domain.OrderRepositoryGetParam{
		ID: param.OrderID,
	})
	if err != nil {
		return nil, err
	}

	orderDto := http.ToOrderResponseDto(order)
	if err := o.enrichOrderItems(ctx, orderDto, order); err != nil {
		return nil, err
	}

	return orderDto, nil
}

func (o *Order) Update(ctx context.Context, param http.UpdateOrderRequestDto) (*http.OrderResponseDto, error) {
	order, err := o.orderRepo.Get(ctx, domain.OrderRepositoryGetParam{
		ID: param.OrderID,
	})
	if err != nil {
		return nil, err
	}

	order.Update(
		param.Data.Address,
		param.Data.Status,
		param.Data.IsPaid,
	)
	err = o.orderRepo.Save(ctx, domain.OrderRepositorySaveParam{
		Order: *order,
	})
	if err != nil {
		return nil, err
	}

	orderDto := http.ToOrderResponseDto(order)
	if err := o.enrichOrderItems(ctx, orderDto, order); err != nil {
		return nil, err
	}

	return orderDto, nil
}
