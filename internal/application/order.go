package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
)

type Order struct {
	vnpaypaymentService VNPayPaymentService
	orderRepo           domain.OrderRepository
	orderService        domain.OrderService
	productRepo         domain.ProductRepository
	productService      domain.ProductService
	cartRepo            domain.CartRepository
}

func ProvideOrder(
	vnpaypaymentService VNPayPaymentService,
	orderRepo domain.OrderRepository,
	orderService domain.OrderService,
	productRepo domain.ProductRepository,
	productService domain.ProductService,
	cartRepo domain.CartRepository,
) *Order {
	return &Order{
		vnpaypaymentService: vnpaypaymentService,
		orderRepo:           orderRepo,
		orderService:        orderService,
		productRepo:         productRepo,
		productService:      productService,
		cartRepo:            cartRepo,
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
		param.Data.RecipientName,
		param.Data.PhoneNumber,
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

	paymentURL, err := o.vnpaypaymentService.GetPaymentURL(ctx, GetPaymentURLVNPayParam{
		ReturnURL: param.Data.ReturnURL,
		Order:     order,
	})
	if err != nil {
		return nil, err
	}
	orderDto := http.ToOrderResponseDto(order, paymentURL)
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
		orderDto := http.ToOrderResponseDto(order, "")
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

	orderDto := http.ToOrderResponseDto(order, "")
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

	orderDto := http.ToOrderResponseDto(order, "")
	if err := o.enrichOrderItems(ctx, orderDto, order); err != nil {
		return nil, err
	}

	return orderDto, nil
}

func (o *Order) VerifyVNPayIPN(ctx context.Context, param http.VerifyVNPayIPNRequestDTO) (*http.VerifyVNPayIPNResponseDTO, error) {
	verifyParam := VerifyIPNVNPayParam{
		Amount:            param.QueryParams.Amount,
		BankTranNo:        param.QueryParams.BankTranNo,
		BankCode:          param.QueryParams.BankCode,
		CardType:          param.QueryParams.CardType,
		OrderInfo:         param.QueryParams.OrderInfo,
		PayDate:           param.QueryParams.PayDate,
		ResponseCode:      param.QueryParams.ResponseCode,
		SecureHash:        param.QueryParams.SecureHash,
		TmnCode:           param.QueryParams.TmnCode,
		TransactionNo:     param.QueryParams.TransactionNo,
		TransactionStatus: param.QueryParams.TransactionStatus,
		TxnRef:            param.QueryParams.TxnRef,
	}
	rspCode, message, err := o.vnpaypaymentService.VerifyIPN(
		ctx,
		verifyParam,
		o.getOrder,
		o.onVerifySuccess,
		o.onVerifyFailure,
	)
	return &http.VerifyVNPayIPNResponseDTO{
		RspCode: rspCode,
		Message: message,
	}, err
}

func (o *Order) getOrder(
	ctx context.Context,
	orderID uuid.UUID,
) (*domain.Order, error) {
	return o.orderRepo.Get(ctx, domain.OrderRepositoryGetParam{
		ID: orderID,
	})
}

func (o *Order) onVerifySuccess(
	ctx context.Context,
	order *domain.Order,
) error {
	order.Update(
		order.Address,
		domain.OrderStatusProcessing,
		true,
	)
	cart, err := o.cartRepo.Get(ctx, domain.CartRepositoryGetParam{
		UserID: order.UserID,
	})
	if err != nil {
		return err
	}
	cart.ClearItems()
	err = o.cartRepo.Save(ctx, domain.CartRepositorySaveParam{
		Cart: *cart,
	})
	if err != nil {
		return err
	}
	productVariantIDs := make([]uuid.UUID, 0, len(order.Items))
	productIDproductVariantIDMap := make(map[uuid.UUID]uuid.UUID)
	for _, item := range order.Items {
		productVariantIDs = append(productVariantIDs, item.ProductVariantID)
		productIDproductVariantIDMap[item.ProductID] = item.ProductVariantID
	}
	products, err := o.productRepo.List(ctx, domain.ProductRepositoryListParam{
		VariantIDs: productVariantIDs,
	})

	productIDproductVariantMap := make(map[uuid.UUID]*domain.ProductVariant)
	for _, p := range *products {
		for _, v := range p.Variants {
			if _, ok := productIDproductVariantIDMap[p.ID]; ok {
				if v.ID == productIDproductVariantIDMap[p.ID] {
					productIDproductVariantMap[p.ID] = &v
				}
			}
		}
	}

	for _, item := range order.Items {
		variant, ok := productIDproductVariantMap[item.ProductID]
		if !ok {
			return domain.ErrNotFound
		}
		variant.DecreaseQuantity(item.Quantity)
	}
	for _, p := range *products {
		_ = o.productRepo.Save(ctx, domain.ProductRepositorySaveParam{
			Product: p,
		})
		if err != nil {
			return err
		}

	}

	return o.orderRepo.Save(ctx, domain.OrderRepositorySaveParam{
		Order: *order,
	})
}

func (o *Order) onVerifyFailure(
	ctx context.Context,
	order *domain.Order,
) error {
	order.Update(
		order.Address,
		domain.OrderStatusCancelled,
		false,
	)
	return o.orderRepo.Save(ctx, domain.OrderRepositorySaveParam{
		Order: *order,
	})
}
