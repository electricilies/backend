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
		param.UserID,
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
	// TODO: map the order to OrderResponseDto, binding the products and variants into dto, then return
	return nil, nil
}

func (o *Order) List(ctx context.Context, param http.ListOrderRequestDto) (*http.PaginationResponseDto[domain.Order], error) {
	// TODO: Get the product and variant and link them to the response dto just like cart
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

	pagination := newPaginationResponseDto(
		*orders,
		*count,
		param.Page,
		param.Limit,
	)
	return pagination, nil
}

func (o *Order) Get(ctx context.Context, param http.GetOrderRequestDto) (*http.OrderResponseDto, error) {
	// TODO: Get the product and variant and link them to the response dto just like cart
	_, err := o.orderRepo.Get(ctx, domain.OrderRepositoryGetParam{
		ID: param.OrderID,
	})
	if err != nil {
		return nil, err
	}
	// TODO: map to response dto
	return nil, nil
}

func (o *Order) Update(ctx context.Context, param http.UpdateOrderRequestDto) (*http.OrderResponseDto, error) {
	// TODO: Get the product and variant and link them to the response dto just like cart
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
	if err != nil {
		return nil, err
	}

	err = o.orderRepo.Save(ctx, domain.OrderRepositorySaveParam{
		Order: *order,
	})
	if err != nil {
		return nil, err
	}
	// TODO: map to response dto
	return nil, nil
}
